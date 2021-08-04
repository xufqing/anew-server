import { DeleteOutlined, PlusOutlined, CodeTwoTone, UsergroupAddOutlined } from '@ant-design/icons';
import { Button, Divider, Modal, message, Menu, Tooltip } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import React, { useState, useRef, useEffect } from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { TableDropdown } from '@ant-design/pro-table';
import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import RecordModal from './components/RecordModal';
import { queryHosts, deleteHost, queryHostGroups } from '@/services/anew/host';
import { useAccess, Access, useModel } from 'umi';

export type optionsType = {
    label: string,
    value: string,
}
const HostList: React.FC = () => {
    const { initialState } = useModel('@@initialState');
    if (!initialState || !initialState.DictObj) {
        return null;
    }
    const authType: optionsType[] = initialState.DictObj.auth_type.map((item: any) => ({ label: item.dict_value, value: item.dict_key }));
    const hostType: optionsType[] = initialState.DictObj.host_type.map((item: any) => ({ label: item.dict_value, value: item.dict_key }));
    const [createVisible, setCreateVisible] = useState<boolean>(false);
    const [updateVisible, setUpdateVisible] = useState<boolean>(false);
    const [recordVisible, setRecordVisible] = useState<boolean>(false);
    const actionRef = useRef<ActionType>();
    const [formValues, setFormValues] = useState<API.HostList>();
    const [hostGroup, setHostGroup] = useState<API.HostGroupList[]>([]);
    const [groupId, setGroupId] = useState<string>();
    const [hostId, setHostId] = useState<number>();
    const access = useAccess();
    const { consoleWin, setConsoleWin } = useModel('global');

    const handleDelete = (record: API.Ids) => {
        if (!record) return;
        if (Array.isArray(record.ids) && !record.ids.length) return;
        const content = `您是否要删除这${Array.isArray(record.ids) ? record.ids.length : ''}项？`;
        Modal.confirm({
            title: '注意',
            content,
            onOk: () => {
                deleteHost(record).then((res) => {
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

    const saveTtys = (val: API.TtyList) => {
        let hosts = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as string)
        if (hosts) {
            hosts.push(val)
        } else {
            hosts = []
            hosts.push(val)
        }
        localStorage.setItem('TABS_TTY_HOSTS', JSON.stringify(hosts));
    }

    useEffect(() => {
        queryHostGroups({ all: true, not_null: true }).then((res) => {
            if (Array.isArray(res.data.data)) {
                setHostGroup([{ id: 0, name: '所有主机' }, ...res.data.data]);
            }
        });
    }, []);

    const columns: ProColumns<API.HostList>[] = [
        {
            title: '主机名',
            dataIndex: 'host_name',
        },
        {
            title: '地址',
            dataIndex: 'ip_address',
        },
        {
            title: '端口',
            dataIndex: 'port',
        },
        {
            title: '主机类型',
            dataIndex: 'host_type',
            valueType: 'select',
            fieldProps: {
                options: hostType,
            },
        },
        {
            title: '认证类型',
            dataIndex: 'auth_type',
            valueType: 'select',
            fieldProps: {
                options: authType,
            },
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
                    <Tooltip title="控制台">
                        <CodeTwoTone
                            style={{ fontSize: '17px', color: 'blue' }}
                            onClick={() => {
                                let actKey = "tty1"
                                let hosts = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as string) || []
                                if (hosts) {
                                    actKey = "tty" + (hosts.length + 1).toString()
                                }
                                const hostsObj: API.TtyList = { hostname: record.host_name, ipaddr: record.ip_address, port: record.port, id: record.id.toString(), actKey: actKey, secKey: null }
                                saveTtys(hostsObj)
                                if (consoleWin) {
                                    if (!consoleWin.closed) {
                                        consoleWin.focus();
                                    } else {
                                        setConsoleWin(window.open('/asset/console', 'consoleTrm'));
                                    }
                                } else {
                                    setConsoleWin(window.open('/asset/console', 'consoleTrm'));
                                }
                            }}
                        />
                    </Tooltip>

                    {/* <Divider type="vertical" />
                <Tooltip title="详情">
                  <FileDoneOutlined
                    style={{ fontSize: '17px', color: '#52c41a' }}
                    onClick={() => {
                      setFormValues(record);
                      handleDescModalVisible(true);
                    }}
                  />
                </Tooltip> */}
                    <Divider type="vertical" />
                    <TableDropdown
                        key="actionGroup"
                        onSelect={(key) => {

                            switch (key) {
                                case 'delete':
                                    handleDelete({ ids: [record.id] });
                                    break;
                                case 'edit':
                                    setFormValues(record);
                                    setUpdateVisible(true);
                                    break;
                                case 'record':
                                    setHostId(record.id)
                                    setRecordVisible(true);
                                    break;
                            }
                        }}
                        menus={[
                            { key: 'edit', name: '修改' },
                            { key: 'record', name: '操作录像' },
                            { key: 'delete', name: '删除' },
                        ]}
                    />
                </>
            ),
        },
    ];

    return (
        <PageHeaderWrapper>
            {/* 权限控制显示内容 */}
            {access.hasPerms(['admin', 'host:list']) && <ProTable
                actionRef={actionRef}
                rowKey="id"
                toolBarRender={(action, { selectedRows }) => [
                    <Access accessible={access.hasPerms(['admin', 'host:create'])}>
                        <Button key="1" type="primary" onClick={() => setCreateVisible(true)}>
                            <PlusOutlined /> 新建
                        </Button>
                    </Access>,
                    selectedRows && selectedRows.length > 0 && (
                        <Access accessible={access.hasPerms(['admin', 'host:delete'])}>
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
                request={async (params) => queryHosts({ params }).then((res) => res.data)}
                params={{ groupId, }}
                columns={columns}
                rowSelection={{}}
                tableRender={(_, dom) => hostGroup.length > 1 ? (
                    <div style={{ display: 'flex', width: '100%', }}>
                        <Menu
                            onSelect={(e) => {
                                if (e.key === '0') {
                                    setGroupId('');
                                } else {
                                    setGroupId(e.key);
                                }
                            }}
                            style={{ width: 156 }}
                            defaultSelectedKeys={['0']}
                            defaultOpenKeys={['sub1']}
                            mode="inline"
                        >
                            <Menu.SubMenu
                                key="sub1"
                                title={
                                    <span>
                                        <UsergroupAddOutlined />
                                        <span>主机分组</span>
                                    </span>
                                }
                            >
                                {hostGroup &&
                                    hostGroup.map((item) => <Menu.Item key={item.id}>{item.name}</Menu.Item>)}
                            </Menu.SubMenu>
                        </Menu>
                        <div style={{ flex: 1, }}>
                            {dom}
                        </div>
                    </div>
                ) : <div style={{ flex: 1, }}>
                    {dom}
                </div>}
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
            {recordVisible && (
                <RecordModal
                    handleChange={setRecordVisible}
                    modalVisible={recordVisible}
                    hostId={hostId}
                />
            )}
        </PageHeaderWrapper>
    );
};

export default HostList;