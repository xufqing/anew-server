import React, { useRef, useEffect, useState } from 'react';
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { WebLinksAddon } from 'xterm-addon-web-links';
import { AttachAddon } from 'xterm-addon-attach';
import { SearchAddon } from 'xterm-addon-search';
import 'xterm/css/xterm.css';
import type { ITerminalOptions } from 'xterm'
import styles from '../index.module.css';

export type SSHTerminalProps = {
    hostId: string;
    arrNum: number;
};
const SSHTerminal: React.FC<SSHTerminalProps> = (props) => {
    const { hostId, arrNum } = props
    const termRf = useRef<HTMLDivElement>(null);
    const [xterm, setXterm] = useState<Terminal>();
    const terminalOpts: ITerminalOptions = {
        allowTransparency: true,
        fontFamily: 'operator mono,SFMono-Regular,Consolas,Liberation Mono,Menlo,monospace',
        fontSize: 15,
        cursorStyle: 'underline',
        cursorBlink: true,
    };
    useEffect(() => {
        const terminal = new Terminal(terminalOpts);
        setXterm(terminal);
    }, []);

    useEffect(() => {
        let webSocket: any;
        const fitPlugin = new FitAddon();
        const handleTerminalInit = async () => {
            if (termRf.current && xterm) {
                const token = localStorage.getItem('token');
                const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
                webSocket = new WebSocket(`${protocol}//${window.location.host}/api/v1/host/ssh?host_id=${hostId}&token=${token}`);
                const webLinksPlugin = new WebLinksAddon();
                const SearchPlugin = new SearchAddon();
                xterm.loadAddon(fitPlugin);
                xterm.loadAddon(webLinksPlugin); //链接检测
                xterm.loadAddon(SearchPlugin);
                // 监听窗口
                webSocket.onopen = (e: any) => {
                    console.log('socket连接成功');
                    const attachPlugin = new AttachAddon(webSocket);
                    xterm.loadAddon(attachPlugin);
                    xterm.open(termRf.current as HTMLDivElement);
                    const _cols = parseInt(String(termRf.current ? termRf.current.clientWidth / 8 : '0'))
                    const _rows = parseInt(String(termRf.current ? termRf.current.clientHeight / 16 : '0'))
                    xterm.resize(_cols, _rows);
                    xterm.focus();
                    xterm.clear();

                };
                //接收服务端消息
                webSocket.onmessage = (e: any) => {
                    const message = e.data;
                    if (message.indexOf) {
                        if (message.indexOf('Anew-Sec-WebSocket-Key') != -1) {
                            let val = JSON.parse(localStorage.getItem('TABS_TTY_HOSTS') as any)
                            let secKey = message.substring(message.lastIndexOf(':') + 1, message.length).replace(/[\r\n]/g, "")
                            val[arrNum].secKey = secKey
                            localStorage.setItem('TABS_TTY_HOSTS', JSON.stringify(val));
                        }
                    }
                };


                //监听关闭事件
                webSocket.onclose = (e: any) => {
                    setTimeout(() => xterm.write('\r\n?[31mConnection is closed.?[0m\r\n'), 500);
                };
                //监听窗口大小
                xterm.onResize((size) => {
                    const resize = JSON.stringify({
                        type: 'resizePty',
                        cols: size.cols,
                        rows: size.rows,
                    });
                    webSocket.send(resize);
                    xterm.resize(size.cols, size.rows);
                });
            }
        };

        handleTerminalInit();
        // 组件销毁前操作
        const closeWs = () => {
            if (webSocket) {
                const closed = JSON.stringify({
                    type: 'closePty',
                    close: true,
                });
                webSocket.send(closed);
            }
        }
        const fitListener = () => {
            fitPlugin.fit()
        };
        window.addEventListener('resize', fitListener)
        return () => {
            closeWs(); //退出ws
            window.removeEventListener('resize', fitListener)
        };
    }, [termRf, xterm]);



    return (
        <div className={styles.container}>
            <div className={styles.terminal}>
                <div ref={termRf} />
            </div>
        </div>
    );
}

export default SSHTerminal;