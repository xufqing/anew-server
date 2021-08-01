import { DeleteOutlined, PlusOutlined, FormOutlined, SafetyCertificateOutlined } from '@ant-design/icons';
import { Button, Tooltip, Divider, Modal, message } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import PermsForm from './components/PermsForm';
import { queryRoles, deleteRole } from '@/services/anew/role';
import { useAccess, Access } from 'umi';

const RoleList: React.FC = () => {

    const [createVisible, setCreateVisible] = useState<boolean>(false);
    const [updateVisible, setUpdateVisible] = useState<boolean>(false);
    const [permsVisible, setPermsVisible] = useState<boolean>(false);
    const actionRef = useRef<ActionType>();
    const [formValues, setFormValues] = useState<API.RoleList>();
    const access = useAccess();

    const handleDelete = (record: API.Ids) => {
        if (!record) return;
        if (Array.isArray(record.ids) && !record.ids.length) return;
        const content = `您是否要删除这${Array.isArray(record.ids) ? record.ids.length : ''}项？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteRole(record).then((res) => {
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

    const columns: ProColumns<API.RoleList>[] = [
        {
            title: '名称',
            dataIndex: 'name',
        },
        {
            title: '关键字',
            dataIndex: 'keyword',
            search: false,
        },
        {
            title: '说明',
            dataIndex: 'desc',
        },
        {
            title: '创建人',
            dataIndex: 'creator',
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => (
                <>
                    <Access accessible={access.hasPerms(['admin', 'role:update:perms'])}>
                        <Tooltip title="设置权限">
                            <SafetyCertificateOutlined
                                style={{ fontSize: '17px', color: 'blue' }}
                                onClick={() => {
                                    if (record.keyword != 'admin') {
                                        setFormValues(record);
                                        setPermsVisible(true);
                                    } else {
                                        message.info("管理员拥有所有权限");
                                    }
                                }}
                            />
                        </Tooltip>
                    </Access>
                    <Divider type="vertical" />
                    <Access accessible={access.hasPerms(['admin', 'role:update'])}>
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
                    <Access accessible={access.hasPerms(['admin', 'role:delete'])}>
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
            {access.hasPerms(['admin', 'role:list']) && <ProTable
                actionRef={actionRef}
                rowKey="id"
                toolBarRender={(action, { selectedRows }) => [
                    <Access accessible={access.hasPerms(['admin', 'role:create'])}>
                        <Button key="1" type="primary" onClick={() => setCreateVisible(true)}>
                            <PlusOutlined /> 新建
                        </Button>
                    </Access>,
                    selectedRows && selectedRows.length > 0 && (
                        <Access accessible={access.hasPerms(['admin', 'role:delete'])}>
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
                request={async (params) => queryRoles({ params }).then((res) => res.data)}
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
            {permsVisible && (
                <PermsForm
                    actionRef={actionRef}
                    handleChange={setPermsVisible}
                    modalVisible={permsVisible}
                    values={formValues}
                />
            )}
        </PageHeaderWrapper>
    );
};

export default RoleList;