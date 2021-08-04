import React, { useState, useEffect } from 'react';
import { CodeOutlined } from '@ant-design/icons';
import { Modal, Tabs, Tree, message } from 'antd';
import { queryHosts, queryHostGroups, queryHostByGroupId } from '@/services/anew/host';


const { TabPane } = Tabs;
const { DirectoryTree } = Tree;

export type FileManagerProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    setTtys: (host: API.TtyList[]) => void;
    setActiveKey: (actKey: string) => void;
};


const ServerManager: React.FC<FileManagerProps> = (props) => {
    const { modalVisible, handleChange, setTtys, setActiveKey } = props;
    const [treeData, setTreeData] = useState();
    const [loadedKeys, setLoadedKeys] = useState<string[]>([]);
    const [selectedKeys, setSelectedKeys] = useState<string[]>([]);
    const [addTty, setAddTty] = useState<API.TtyList>();

    const updateTreeData = (list: any, key: React.Key, children: any): any => {
        return list.map((node: any) => {
            if (node.key === key) {
                return { ...node, children };
            }

            if (node.children) {
                return { ...node, children: updateTreeData(node.children, key, children) };
            }
            return node;
        })
    };
    const onLoadData = ({ key, children }: any) => {
        return new Promise<void>(resolve => {
            if (children) {
                resolve();
                return;
            }
            if (key === "0") {
                queryHosts().then((res) => {
                    const selectHosts = Array.isArray(res.data.data) ? res.data.data.map((item: any) => ({
                        ...item,
                        title: `${item.host_name} (${item.ip_address}:${item.port})`,
                        key: `key-${key}-${item.id}`,
                        isLeaf: true,
                        icon: <CodeOutlined />,
                    })) : null
                    setTreeData((origin) => updateTreeData(origin, key, selectHosts))
                }
                )
            } else {
                queryHostByGroupId(key).then((res) => {
                    const selectHosts = Array.isArray(res.data.data) ? res.data.data.map((item: any) => ({
                        ...item,
                        title: `${item.host_name} (${item.ip_address}:${item.port})`,
                        key: `key-${key}-${item.id}`,
                        isLeaf: true,
                        icon: <CodeOutlined />,
                    })) : null
                    setTreeData((origin) => updateTreeData(origin, key, selectHosts))
                }
                )
            }
            setLoadedKeys([...loadedKeys, key])
            resolve();
        });
    }

    const onSelect = (selectedKeysValue: any, info: any) => {
        setSelectedKeys(selectedKeysValue);
        const key = selectedKeysValue[0].split("-")
        if (key[0] === "key") {
            const hostsObj = { hostname: info.node.host_name, ipaddr: info.node.ip_address, port: info.node.port, id: info.node.id.toString(), actKey: "tty1", secKey: null }
            setAddTty(hostsObj)
        }
    };

    const getHostDir = () => {
        queryHostGroups({ all: true, not_null: true }).then((res) => {
            if (Array.isArray(res.data.data)) {
                var resp = [{ id: 0, name: '所有主机' }, ...res.data.data]
                resp = resp.map(({ id, name, ...item }) => ({
                    ...item,
                    title: name,
                    key: `${id}`,
                }))
                setLoadedKeys([]);
                setTreeData(resp as any);
            }
        });
    }
    const tabcallback = (key: string) => {
        switch (key) {
            case "1":
                getHostDir();
        }
    }
    const handleOk = () => {
        const key = selectedKeys[0].split("-");
        if (key[0] === "key") {
            let hosts = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as string) || [];
            if (hosts) {
                let newAddtty = {};
                Object.assign(newAddtty, addTty);
                (newAddtty as API.TtyList).actKey = "tty" + (hosts.length + 1).toString();
                hosts.push(newAddtty);
                setTtys(hosts);
                localStorage.setItem('TABS_TTY_HOSTS', JSON.stringify(hosts));
                setActiveKey((newAddtty as API.TtyList).actKey);
            }
            handleChange(false);
        } else {
            message.info("未选择服务器");
        }
    }
    useEffect(() => {
        // 默认加载主机组目录
        getHostDir();
    }, []);
    return (
        <Modal
            title="选择服务器"
            visible={modalVisible}
            onCancel={() => handleChange(false)}
            onOk={handleOk}
            okText="确认"
            cancelText="取消"
        >
            <Tabs defaultActiveKey="1" onChange={tabcallback} tabPosition="left">
                <TabPane tab="主机组" key="1">
                    <DirectoryTree loadData={onLoadData} treeData={treeData} loadedKeys={loadedKeys} selectedKeys={selectedKeys} onSelect={onSelect} />
                </TabPane>
                <TabPane tab="项目组" key="2">
                    项目分组
                </TabPane>
            </Tabs>

        </Modal>
    );
}

export default ServerManager;