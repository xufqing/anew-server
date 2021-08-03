import {
    FolderOutlined,
    DesktopOutlined
} from '@ant-design/icons';
import React, { useEffect, useState } from 'react';
import { message, Tabs, Button, Tooltip } from 'antd';
import DraggableTabs from '@/components/DraggableTabs';
import Terminal from './components/Terminal';
import FileManager from './components/FileManager';
import ServerManager from './components/ServerManager';

const { TabPane } = Tabs;
const Console: React.FC = () => {
    const [Ttys, setTtys] = useState<API.TtyList[]>(JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as any));
    const [activeKey, setActiveKey] = useState<string>('tty1');
    const [fileVisible, setFileVisible] = useState(false);
    const [serverVisible, setServerVisible] = useState(false);
    const [webSocketKey, setWebSocketKey] = useState<string>('');
    const callback = (key: string) => {
        setActiveKey(key)
    }

    const onEdit = (targetKey: any, action: any) => {
        if (action === "remove") {
            remove(targetKey)
        }
        //console.log(action, targetKey)
    };
    const remove = (key: React.Key) => {
        let val = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as any)
        const index = val.map((item: API.TtyList) => item.actKey).indexOf(key)
        val.splice(index, 1)
        localStorage.setItem('TABS_TTY_HOSTS', JSON.stringify(val));
        setTtys(val)
        let [last] = [...val].reverse()
        if (typeof (last) !== 'undefined') {
            setActiveKey(last.actKey)
        }
    }
    useEffect(() => {
        document.title = '控制台管理器'
    }, []);

    //监听事件
    useEffect(() => {
        window.onbeforeunload = () => {
            localStorage.removeItem('TABS_TTY_HOSTS');
        }
        const refetch = (e: any) => {
            if (e.key === "TABS_TTY_HOSTS") {
                let val = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as any)
                setTtys(val)
                if (val) {
                    let [last] = [...val].reverse()
                    setActiveKey(last.actKey)
                }
            }
        }
        window.addEventListener('storage', refetch)
        return () => {
            window.removeEventListener('storage', refetch)
        };
    }, [Ttys]);
    const operations = (
        <div>
            <Button icon={<DesktopOutlined />} onClick={() => {
                setServerVisible(true)
            }} >我的服务器</Button>
            <Button icon={<FolderOutlined />} onClick={() => {
                const val = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as any)
                if (val) {
                    const num = val.map((item: API.TtyList) => item.actKey).indexOf(activeKey)
                    if (num !== -1) {
                        setWebSocketKey(val[num].secKey)
                        setFileVisible(true)
                    } else {
                        message.error("未连接")
                    }
                } else {
                    message.error("未连接")
                }
            }}>文件管理器</Button>
        </div>
    )

    return (
        <div>
            <DraggableTabs
                size="small"
                type="editable-card"
                tabBarStyle={{ margin: 0 }}
                hideAdd
                tabBarExtraContent={operations}
                onChange={callback}
                activeKey={activeKey}
                onEdit={onEdit}>
                {Ttys && Ttys.map((v, i) => (
                    <TabPane forceRender tab={<Tooltip title={v.ipaddr + `:` + v.port}>{v.hostname}</Tooltip>} key={v.actKey}>
                        <Terminal hostId={v.id} arrNum={i} />
                    </TabPane>
                ))}
            </DraggableTabs>
            {serverVisible && (
                <ServerManager
                    handleChange={setServerVisible}
                    setTtys={setTtys}
                    setActiveKey={setActiveKey}
                    modalVisible={serverVisible}
                />
            )}
            {fileVisible && (
                <FileManager
                    handleChange={setFileVisible}
                    connectId={encodeURIComponent(webSocketKey)}
                    modalVisible={fileVisible}
                />
            )}
        </div>
    );
}

export default Console;