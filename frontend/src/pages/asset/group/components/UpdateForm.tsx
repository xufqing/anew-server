import React, { useEffect, useState } from 'react';
import { queryHosts, updateHostGroup } from '@/services/anew/host';
import ProForm, { ModalForm, ProFormText } from '@ant-design/pro-form';
import { message, Transfer, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';
import type { TransferItem } from 'antd/lib/transfer'

export type UpdateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.HostGroupList | undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange, values } = props;
    const [hostData, setHostData] = useState<TransferItem[]>([]);
    const [targetKeys, setTargetKeys] = useState<string[]>([]);

    const transferChange = (keys: string[]) => {
        setTargetKeys(keys);
    };

    useEffect(() => {
        console.log(hostData)
    },[hostData]);
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
        transferChange(values?.hosts_id as unknown as string[])
    }, []);

    return (
        <ModalForm
            title="更新主机组"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                updateHostGroup(v as API.HostGroupParams, values?.id).then((res) => {
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
                <ProFormText
                    name="name"
                    label="名称"
                    width="md"
                    initialValue={values?.name}
                    rules={[{ required: true }]}
                />
                <ProFormText name="desc" label="说明" width="md" initialValue={values?.desc} />
                <Form.Item label="选择主机" name="hosts">
                    <Transfer
                        dataSource={hostData}
                        showSearch
                        listStyle={{
                            width: 320,
                            height: 280,
                        }}
                        //operations={['加入', '退出']}
                        targetKeys={targetKeys ? targetKeys : []}
                        onChange={transferChange}
                        render={(item) => `${item.title}(${item.description})`}
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default UpdateForm;