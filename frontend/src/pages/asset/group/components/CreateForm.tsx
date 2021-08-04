import React, { useEffect, useState } from 'react';
import { queryHosts, createHostGroup } from '@/services/anew/host';
import ProForm, { ModalForm, ProFormText } from '@ant-design/pro-form';
import { message, Transfer, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';
import type { TransferItem } from 'antd/lib/transfer'

export type CreateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
};

const CreateForm: React.FC<CreateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange } = props;
    const [hostData, setHostData] = useState<TransferItem[]>([]);
    const [targetKeys, setTargetKeys] = useState<string[]>([]);

    const transferChange = (keys: string[]) => {
        setTargetKeys(keys);
    };

    useEffect(() => {
        queryHosts({ all: true }).then((res) => {
            if (Array.isArray(res.data.data)) {
                setHostData(
                    res.data.data.map((item: API.HostList) => ({
                        key: item.id,
                        title: item.host_name,
                        description: item.ip_address,
                    })),
                );
            }
        });
    }, []);

    return (
        <ModalForm
            title="新建主机组"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                createHostGroup(v as API.HostGroupParams).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        if (actionRef.current) {
                            actionRef.current.reload();
                        }
                    }
                });
                return true;
            }}
        >
            <ProForm.Group>
                <ProFormText name="name" label="名称" width="md" rules={[{ required: true }]} />
                <ProFormText name="desc" label="说明" width="md" />
                <Form.Item label="选择主机" name="hosts">
                    <Transfer
                        dataSource={hostData}
                        showSearch
                        listStyle={{
                            width: 320,
                            height: 280,
                        }}
                        //operations={['加入', '退出']}
                        targetKeys={targetKeys}
                        onChange={transferChange}
                        render={(item) => `${item.title}(${item.description})`}
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;