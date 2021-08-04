import { DeleteOutlined } from '@ant-design/icons';
import { Tooltip, Divider, Modal, message } from 'antd';
import React, { useRef } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { querySessions, deleteSession } from '@/services/anew/session';
import { useAccess, Access } from 'umi';

const SessionList: React.FC = () => {
    const actionRef = useRef<ActionType>();
    const access = useAccess();

    const handleDelete = (record: any) => {
        if (!record) return;
        const content = `您是否要注销该连接？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteSession(record).then((res) => {
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

    const columns: ProColumns<API.SessionList>[] = [
        {
            dataIndex: 'index',
            valueType: 'indexBorder',
            width: 48,
        },
        {
            title: '用户名',
            dataIndex: 'user_name',
        },
        {
            title: '主机名',
            dataIndex: 'host_name',
        },
        {
            title: '接入时间',
            dataIndex: 'connect_time',
            //sorter: true,
        },
        {
            title: '标识',
            dataIndex: 'connect_id',
        },
        {
            title: '操作',
            dataIndex: 'option',
            valueType: 'option',
            render: (_, record) => (
                <>
                    <Divider type="vertical" />
                    <Access accessible={access.hasPerms(['admin', 'session:delete'])}>
                        <Tooltip title="注销">
                            <DeleteOutlined
                                style={{ fontSize: '17px', color: 'red' }}
                                onClick={() => handleDelete({ key: record.connect_id })}
                            />
                        </Tooltip>
                    </Access>
                </>
            ),
        },
    ];

    return (
        <PageHeaderWrapper>
            {access.hasPerms(['admin', 'session:list']) && <ProTable
                pagination={false}
                search={false}
                actionRef={actionRef}
                rowKey="connect_id"
                request={(params) => querySessions({ ...params })}
                columns={columns}
            />}
        </PageHeaderWrapper>
    );
};

export default SessionList;
