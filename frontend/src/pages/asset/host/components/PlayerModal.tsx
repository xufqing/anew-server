import { Modal } from 'antd';
import React, { useEffect } from 'react';
import { useScript } from '@/utils/useScript';

export type RecordModalProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    values: API.RecordList | undefined;
};

const PlayerModal: React.FC<RecordModalProps> = (props) => {
    const { modalVisible, handleChange, values } = props;
    const [loaded, error] = useScript('/asciinema-player.js');

    const token = `/api/v1/host/record/play?record=${values?.connect_id}&token=${localStorage.getItem('token')}`
    const head = `<head><link rel="stylesheet" type="text/css" href="/asciinema-player.css" /></head>`;
    const defaultprops = ` preload`
    const body = `<body><asciinema-player src="${token}" ${defaultprops} ></asciinema-player></body>`;

    useEffect(() => {
        if (!loaded) return;
    }, [loaded, error]);

    return (
        <Modal
            title="播放录像"
            visible={modalVisible}
            onCancel={() => handleChange(false)}
            footer={null}
            style={{ top: 0 }}
            width={'88%'}
        >
            <div style={{ flex: 1,  }} dangerouslySetInnerHTML={{ __html: head + body }} />

        </Modal>
    );
};

export default PlayerModal;
