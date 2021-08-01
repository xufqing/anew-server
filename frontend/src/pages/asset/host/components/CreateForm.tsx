import React, { useState } from 'react';
import { createHost } from '@/services/anew/host';
import ProForm, { ModalForm, ProFormText, ProFormSelect } from '@ant-design/pro-form';
import { message } from 'antd';
import type { ActionType } from '@ant-design/pro-table';
import type { optionsType } from '../index';


export type CreateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    authType: optionsType[];
    hostType: optionsType[];
    actionRef: React.MutableRefObject<ActionType | undefined>;
};

const CreateForm: React.FC<CreateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange, authType, hostType } = props;
    const [isKey, setIsKey] = useState<boolean>(false);

    return (
        <ModalForm
            title="新建主机"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                createHost(v as any).then((res) => {
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
                <ProFormText name="host_name" label="主机名称" width="md" />
                <ProFormText name="ip_address" label="地址" width="md" rules={[{ required: true }]} />
                <ProFormText name="port" label="端口" width="md" rules={[{ required: true }]} />
                <ProFormSelect
                    name="host_type"
                    label="主机类型"
                    width="md"
                    hasFeedback
                    options={hostType}
                //onChange
                />
                <ProFormSelect
                    name="auth_type"
                    label="认证类型"
                    hasFeedback
                    width="md"
                    options={authType}
                    rules={[{ required: true, message: '请选择认证类型' }]}
                    fieldProps={{
                        onSelect: (e) => {
                            if (e === 'key') {
                                setIsKey(true);
                            } else {
                                setIsKey(false);
                            }
                        },
                    }}
                />
                <ProFormText name="user" label="用户" width="md" rules={[{ required: true }]} />

                {isKey ? (
                    <ProFormText name="privatekey" label="密钥路径" width="md" rules={[{ required: true }]} />
                ) : null}
                {isKey ? (
                    <ProFormText.Password
                        label="密钥密码"
                        name="key_passphrase"
                        width="md"
                        placeholder="如果有密码"
                    />
                ) : (
                    <ProFormText.Password
                        label="认证密码"
                        name="password"
                        width="md"
                        rules={[{ required: true }]}
                    />
                )}
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;