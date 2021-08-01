import { DeleteOutlined, PlusOutlined, FormOutlined } from '@ant-design/icons';
import { Button, Tooltip, Divider, Modal, message } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import { queryApis, deleteApi } from '@/services/anew/api';
import { useAccess, Access } from 'umi';

const ApiList: React.FC = () => {
    const [createVisible, setCreateVisible] = useState<boolean>(false);
    const [updateVisible, setUpdateVisible] = useState<boolean>(false);
    const actionRef = useRef<ActionType>();
    const [formValues, setFormValues] = useState<API.ApiList>();
    const access = useAccess();

    const handleDelete = (record: API.Ids) => {
        if (!record) return;
        if (Array.isArray(record.ids) && !record.ids.length) return;
        const content = `您是否要删除这${Array.isArray(record.ids) ? record.ids.length : ''}项？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteApi(record).then((res) => {
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

    const columns: ProColumns<API.ApiList>[] = [
        {
            title: '名称',
            dataIndex: 'name',
        },
        {
            title: '请求方式',
            dataIndex: 'method',
            // render: (_, row) => {
            //   let color = 'blue';
            //   if (row.method == 'POST') {
            //     color = 'gold';
            //   } else if (row.method == 'PATCH') {
            //     color = 'lime';
            //   } else if (row.method == 'PUT') {
            //     color = 'green';
            //   } else if (row.method == 'DELETE') {
            //     color = 'red';
            //   }
            //   return <Tag color={color}>{row.method}</Tag>;
            // },
        },
        {
            title: '访问路径',
            dataIndex: 'path',
        },
        {
            title: '权限标识',
            dataIndex: 'perms_tag',
        },
        {
            title: '说明',
            dataIndex: 'desc',
            search: false,
        },
        {
            title: '创建人',
            dataIndex: 'creator',
            search: false,
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record: API.ApiList) => (
                <>
                    <Access accessible={access.hasPerms(['admin', 'api:update'])}>
                        <Tooltip title="修改">
                            <FormOutlined
                                style={{ fontSize: '17px', color: '#52c41a' }}
                                onClick={() => {
                                    setFormValues(record);
                                    setUpdateVisible(true);
                                }}
                            />
                        </Tooltip>
                    </Access>
                    <Divider type="vertical" />
                    <Access accessible={access.hasPerms(['admin', 'api:delete'])}>
                        <Tooltip title="删除">
                            <DeleteOutlined
                                style={{ fontSize: '17px', color: 'red' }}
                                onClick={() => handleDelete({ ids: [record.id] })}
                            />
                        </Tooltip>
                    </Access>
                </>
            ),
        },
    ];

    return (
        <PageHeaderWrapper>
            {/* 权限控制显示内容 */}
            {access.hasPerms(['admin', 'api:list']) && <ProTable
                actionRef={actionRef}
                rowKey="id"
                pagination={false}
                search={false}
                toolBarRender={(action, { selectedRows }) => [
                    <Access accessible={access.hasPerms(['admin', 'api:create'])}>
                        <Button key="1" type="primary" onClick={() => setCreateVisible(true)}>
                            <PlusOutlined /> 新建
                        </Button>
                    </Access>,
                    selectedRows && selectedRows.length > 0 && (
                        <Access accessible={access.hasPerms(['admin', 'api:delete'])}>
                            <Button
                                key="2"
                                type="primary"
                                onClick={() => handleDelete({ ids: selectedRows.map((item) => item.id) })}
                                danger
                            >
                                <DeleteOutlined /> 删除
                            </Button>
                        </Access>
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
                request={async (params) => queryApis({ ...params })}
                columns={columns}
                rowSelection={{}}
            />}
            {createVisible && (
                <CreateForm
                    actionRef={actionRef}
                    handleChange={setCreateVisible}
                    modalVisible={createVisible}
                />
            )}
            {updateVisible && (
                <UpdateForm
                    actionRef={actionRef}
                    handleChange={setUpdateVisible}
                    modalVisible={updateVisible}
                    values={formValues}
                />
            )}
        </PageHeaderWrapper>
    );
};

export default ApiList;