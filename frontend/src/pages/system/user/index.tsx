import { DeleteOutlined, PlusOutlined, FormOutlined } from '@ant-design/icons';
import { Button, Tooltip, Divider, Modal, message } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import { queryUsers, deleteUser } from '@/services/anew/user';
import { useAccess, Access } from 'umi';

const UserList: React.FC = () => {

    const [createVisible, setCreateVisible] = useState<boolean>(false);
    const [updateVisible, setUpdateVisible] = useState<boolean>(false);
    const actionRef = useRef<ActionType>();
    const [formValues, setFormValues] = useState<API.UserList>();
    const access = useAccess();

    const handleDelete = (record: API.Ids) => {
        if (!record) return;
        if (Array.isArray(record.ids) && !record.ids.length) return;
        const content = `您是否要删除这${Array.isArray(record.ids) ? record.ids.length : ''}项？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteUser(record).then((res) => {
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

    const columns: ProColumns<API.UserList>[] = [
        {
            title: '用户名',
            dataIndex: 'username',
        },
        {
            title: '姓名',
            dataIndex: 'name',
        },
        {
            title: '手机',
            dataIndex: 'mobile',
        },
        {
            title: '邮箱',
            dataIndex: 'email',
            search: false,
        },
        {
            title: '角色',
            dataIndex: 'role',
            search: false,
            render: (_, record: API.UserList) => {
                return record.role.name;
            },
        },
        {
            title: '部门',
            dataIndex: 'dept',
            search: false,
            render: (_, record: API.UserList) => {
                return record.dept.name;
            },
        },
        {
            title: '创建人',
            dataIndex: 'creator',
        },
        {
            title: '状态',
            dataIndex: 'status',
            valueEnum: {
                true: {
                    text: '激活',
                    status: 'Processing',
                },
                false: {
                    text: '禁用',
                    status: 'Error',
                },
            },
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record: API.UserList) => (
                <>
                    <Access accessible={access.hasPerms(['admin', 'user:update'])}>
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
                    <Access accessible={access.hasPerms(['admin', 'user:delete'])}>
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
            {access.hasPerms(['admin', 'user:list']) && <ProTable
                actionRef={actionRef}
                rowKey="id"
                toolBarRender={(action, { selectedRows }) => [
                    <Access accessible={access.hasPerms(['admin', 'user:create'])}>
                        <Button key="1" type="primary" onClick={() => setCreateVisible(true)}>
                            <PlusOutlined /> 新建
                        </Button>
                    </Access>,
                    selectedRows && selectedRows.length > 0 && (
                        <Access accessible={access.hasPerms(['admin', 'user:delete'])}>
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
                request={async (params) => queryUsers({ params }).then((res) => res.data)}
                columns={columns}
                rowSelection={{}}
            />}
            {createVisible && (
                <CreateForm
                    actionRef={actionRef}
                    onChange={setCreateVisible}
                    modalVisible={createVisible}
                />
            )}
            {updateVisible && (
                <UpdateForm
                    actionRef={actionRef}
                    onChange={setUpdateVisible}
                    modalVisible={updateVisible}
                    values={formValues}
                />
            )}
        </PageHeaderWrapper>
    );
};

export default UserList;