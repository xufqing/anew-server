import React, { useState, useRef } from 'react';
import { DeleteOutlined, VideoCameraOutlined } from '@ant-design/icons';
import { queryRecords, deleteRecord } from '@/services/anew/record';
import ProTable from '@ant-design/pro-table';
import PlayerModal from './PlayerModal';
import { Drawer, Button, Tooltip, Divider, Modal, message } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';

export type RecordModalProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    hostId: number | undefined;
};

const RecordModal: React.FC<RecordModalProps> = (props) => {
    const { modalVisible, handleChange, hostId } = props;
    const actionRef = useRef<ActionType>();
    const [values, setValues] = useState<API.RecordList>();
    const [playerVisible, setPlayerVisible] = useState<boolean>(false);

    const handleDelete = (record: API.Ids) => {
        if (!record) return;
        if (Array.isArray(record.ids) && !record.ids.length) return;
        const content = `您是否要删除这${Array.isArray(record.ids) ? record.ids.length : ''}项？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteRecord(record).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        if (actionRef.current) {
                            actionRef.current.reload();
                        }
                    }
                });
            },
            onCancel() { },
        });
    };

    const columns: ProColumns<API.RecordList>[] = [
        {
            title: '用户名',
            dataIndex: 'user_name',
        },
        {
            title: '主机名',
            dataIndex: 'host_name',
            search: false,
        },
        {
            title: '标识',
            dataIndex: 'connect_id',
            search: false,
        },
        {
            title: '接入时间',
            dataIndex: 'connect_time',
            valueType: 'dateTime',
            search: false,
        },
        {
            title: '注销时间',
            dataIndex: 'logout_time',
            valueType: 'dateTime',
            search: false,
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => (
                <>
                    <Tooltip title="播放录像">
                        <VideoCameraOutlined
                            style={{ fontSize: '17px', color: '#52c41a' }}
                            onClick={() => {
                                setValues(record);
                                setPlayerVisible(true);
                            }}
                        />
                    </Tooltip>
                    <Divider type="vertical" />
                    <Tooltip title="删除">
                        <DeleteOutlined
                            style={{ fontSize: '17px', color: 'red' }}
                            onClick={() => handleDelete({ ids: [record.id] })}
                        />
                    </Tooltip>
                </>
            ),
        },
    ];

    return (
        <Drawer
            title="操作录像"
            placement="right"
            width={1000}
            visible={modalVisible}
            onClose={() => handleChange(false)}
        >
            <ProTable
                actionRef={actionRef}
                rowKey="connect_time"
                toolBarRender={(action, { selectedRows }) => [
                    selectedRows && selectedRows.length > 0 && (
                        <Button
                            key="1"
                            type="primary"
                            onClick={() => handleDelete({ ids: selectedRows.map((item) => item.id) })}
                            danger
                        >
                            <DeleteOutlined /> 删除
                        </Button>
                    ),
                ]}
                tableAlertRender={({ selectedRowKeys, selectedRows }) => (
                    <div>
                        已选择{' '}
                        <a
                            style={{
                                fontWeight: 600,
                            }}
                        >
                            {selectedRowKeys.length}
                        </a>{' '}
                        项&nbsp;&nbsp;
                    </div>
                )}
                request={(params) => queryRecords(hostId, { params }).then((res) => res.data)}
                columns={columns}
                rowSelection={{}}
            />
            {playerVisible && (
                <PlayerModal
                    handleChange={setPlayerVisible}
                    modalVisible={playerVisible}
                    values={values}
                />
            )}
        </Drawer>
    );
};

export default RecordModal;