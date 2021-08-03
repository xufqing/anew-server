import React, { useState, useEffect } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import {
    FileOutlined,
    FolderTwoTone,
    HomeOutlined,
    ApartmentOutlined,
    DeleteOutlined,
    UploadOutlined,
    LinkOutlined,
    DownloadOutlined,
    InboxOutlined,
} from '@ant-design/icons';
import {
    Drawer,
    Breadcrumb,
    Upload,
    Switch,
    Divider,
    Modal,
    Button,
    Tooltip,
    message,
} from 'antd';
import styles from '../index.module.css';
import ProTable from '@ant-design/pro-table';
import { querySSHFile, deleteSSHFile } from '@/services/anew/ssh';
import lds from 'lodash';

export type FileManagerProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    connectId: string;
};

const FileManager: React.FC<FileManagerProps> = (props) => {
    const { modalVisible, handleChange, connectId } = props;
    const [columnData, setColumnData] = useState<API.SSHFileList[]>([]);
    const [showHidden, setShowHidden] = useState(false);
    const [childrenDrawer, setChildrenDrawer] = useState(false);
    const [currentPathArr, setCurrentPathArr] = useState<string[]>([]);
    const [initPath, setInitPath] = useState<string>();

    const _dirSort = (item: API.SSHFileList) => {
        return item.isDir;
    };

    const getFileData = (key: string, path: string | undefined) => {
        if (!path) return;
        querySSHFile(key, path).then((res) => {
            const obj = lds.orderBy(res.data, [_dirSort, 'name'], ['desc', 'asc']);
            showHidden ? setColumnData(obj) : setColumnData(obj.filter((x) => !x.name.startsWith('.')));
            try {
                // 获取服务器的当前路径
                let pathb = obj[0].path;
                const index = pathb.lastIndexOf('/');
                pathb = pathb.substring(0, index + 1);
                setCurrentPathArr(pathb.split('/').filter((x: any) => x !== ''));
                setInitPath(pathb); // 保存当前路径，刷新用
            } catch (exception) {
                setCurrentPathArr(path.split('/').filter((x) => x !== ''));
                setInitPath(path);
            }
        });
    };

    const getChdirDirData = (key: string, path: string) => {
        const index = currentPathArr.indexOf(path);
        const currentDir = '/' + currentPathArr.splice(0, index + 1).join('/');
        getFileData(key, currentDir);
    };

    const handleDelete = (key: string, path: string) => {
        if (!path) return;
        const index = path.lastIndexOf('/');
        const currentDir = path.substring(0, index + 1);
        const currentFile = path.substring(index + 1, path.length);
        const content = `您是否要删除 ${currentFile}？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteSSHFile(key, path).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        getFileData(key, currentDir);
                    }
                });
            },
            onCancel() { },
        });
    };

    const handleDownload = (key: string, path: string) => {
        if (!path) return;
        const index = path.lastIndexOf('/');
        const currentFile = path.substring(index + 1, path.length);
        const content = `您是否要下载 ${currentFile}？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                const token = localStorage.getItem('token');
                const link = document.createElement('a');
                link.href = `/api/v1/host/ssh/download?key=${key}&path=${path}&token=${token}`;
                document.body.appendChild(link);
                const evt = document.createEvent('MouseEvents');
                evt.initEvent('click', false, false);
                link.dispatchEvent(evt);
                document.body.removeChild(link);
            },
            onCancel() { },
        });
    };

    const uploadProps = {
        name: 'file',
        action: `/api/v1/host/ssh/upload?key=${connectId}&path=${initPath}`,
        multiple: true,
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        // showUploadList: {
        //   removeIcon: false,
        //   showRemoveIcon: false,
        // },
        onChange(info: any) {
            // if (info.file.status !== 'uploading') {
            //   console.log(info.file, info.fileList);
            // }
            //console.log(info);
            if (info.file.status === 'done') {
                message.success(`${info.file.name} file uploaded successfully`);
                getFileData(connectId, initPath); // 刷新数据
            } else if (info.file.status === 'error') {
                message.error(`${info.file.name} file upload failed.`);
            }
        },
        progress: {
            strokeColor: {
                '0%': '#108ee9',
                '100%': '#87d068',
            },
            strokeWidth: 3,
            format: (percent: any) => `${parseFloat(percent.toFixed(2))}%`,
        },
    };


    const columns: ProColumns<API.SSHFileList>[] = [
        {
            title: '名称',
            dataIndex: 'name',
            render: (_, record) =>
                record.isDir ? (
                    <div onClick={() => getFileData(connectId, record.path)} style={{ cursor: 'pointer' }}>
                        <FolderTwoTone />
                        <span style={{ color: '#1890ff', paddingLeft: 5 }}>{record.name}</span>
                    </div>
                ) : (
                    <React.Fragment>
                        {record.isLink ? (
                            <div>
                                <LinkOutlined />
                                <Tooltip title="Is Link">
                                    <span style={{ color: '#3cb371', paddingLeft: 5 }}>{record.name}</span>
                                </Tooltip>
                            </div>
                        ) : (
                            <div>
                                <FileOutlined />
                                <span style={{ paddingLeft: 5 }}>{record.name}</span>
                            </div>
                        )}
                    </React.Fragment>
                ),
        },
        {
            title: '大小',
            dataIndex: 'size',
        },
        {
            title: '修改时间',
            dataIndex: 'mtime',
        },
        {
            title: '属性',
            dataIndex: 'mode',
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) =>
                !record.isDir && !record.isLink ? (
                    <>
                        <Tooltip title="下载文件">
                            <DownloadOutlined
                                style={{ fontSize: '17px', color: 'blue' }}
                                onClick={() => handleDownload(connectId, record.path)}
                            />
                        </Tooltip>
                        <Divider type="vertical" />
                        <Tooltip title="删除文件">
                            <DeleteOutlined
                                style={{ fontSize: '17px', color: 'red' }}
                                onClick={() => handleDelete(connectId, record.path)}
                            />
                        </Tooltip>
                    </>
                ) : null,
        },
    ];


    useEffect(() => {
        getFileData(connectId, '');
    }, []);

    useEffect(() => {
        // 是否显示隐藏文件
        getFileData(connectId, initPath); // 刷新数据
    }, [showHidden]);

    const { Dragger } = Upload;
    return (
        <Drawer
            title="文件管理器"
            placement="right"
            width={800}
            visible={modalVisible}
            onClose={()=>handleChange(false)}
            getContainer={false}
        >
            {/* <input style={{ display: 'none' }} type="file" ref={(ref) => (this.input = ref)} /> */}
            <div className={styles.drawerHeader}>
                <Breadcrumb>
                    <Breadcrumb.Item href="#" onClick={() => getFileData(connectId, '/')}>
                        <ApartmentOutlined />
                    </Breadcrumb.Item>
                    <Breadcrumb.Item href="#" onClick={() => getFileData(connectId, '')}>
                        <HomeOutlined />
                    </Breadcrumb.Item>
                    {currentPathArr.map((item) => (
                        <Breadcrumb.Item key={item} href="#" onClick={() => getChdirDirData(connectId, item)}>
                            <span>{item}</span>
                        </Breadcrumb.Item>
                    ))}
                </Breadcrumb>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                    <span>显示隐藏文件：</span>
                    <Switch
                        checked={showHidden}
                        checkedChildren="开启"
                        unCheckedChildren="关闭"
                        onChange={(v) => {
                            setShowHidden(v);
                        }}
                    />

                    <Button
                        style={{ marginLeft: 10 }}
                        size="small"
                        type="primary"
                        icon={<UploadOutlined />}
                        onClick={() => setChildrenDrawer(true)}
                    >
                        上传文件
                    </Button>
                </div>
            </div>
            <Drawer
                title="上传文件"
                width={320}
                closable={false}
                onClose={() => setChildrenDrawer(false)}
                visible={childrenDrawer}
            >
                <div style={{ height: 150 }}>
                    <Dragger {...uploadProps}>
                        <p className="ant-upload-drag-icon">
                            <InboxOutlined />
                        </p>
                        <p className="ant-upload-text">单击或拖入上传</p>
                        <p className="ant-upload-hint">支持多文件</p>
                    </Dragger>
                </div>
            </Drawer>
            <ProTable
                pagination={false}
                search={false}
                toolBarRender={false}
                rowKey="name"
                dataSource={columnData}
                columns={columns}
            />
        </Drawer>
    );
};
export default FileManager;
