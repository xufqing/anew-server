import React, { useState } from 'react';
import { updateHost } from '@/services/anew/host';
import ProForm, { ModalForm, ProFormText, ProFormSelect } from '@ant-design/pro-form';
import { message } from 'antd';
import type { ActionType } from '@ant-design/pro-table';
import type { optionsType } from '../index';

export type UpdateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    authType: optionsType[];
    hostType: optionsType[];
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.HostList | undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange, values, authType, hostType } = props;
    const [isKey, setIsKey] = useState<boolean>(false);

    return (
        <ModalForm
            title="更新主机"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                updateHost(v as any, values?.id).then((res) => {
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
                <ProFormText name="host_name" label="主机名称" width="md" initialValue={values?.host_name} />
                <ProFormText
                    name="ip_address"
                    label="地址"
                    width="md"
                    rules={[{ required: true }]}
                    initialValue={values?.ip_address}
                />
                <ProFormText
                    name="port"
                    label="端口"
                    width="md"
                    rules={[{ required: true }]}
                    initialValue={values?.port}
                />
                <ProFormSelect
                    name="host_type"
                    label="主机类型"
                    width="md"
                    hasFeedback
                    options={hostType}
                    initialValue={values?.host_type}
                />
                <ProFormSelect
                    name="auth_type"
                    label="认证类型"
                    hasFeedback
                    width="md"
                    initialValue={values?.auth_type}
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

                <ProFormText
                    name="user"
                    label="用户"
                    width="md"
                    rules={[{ required: true }]}
                    initialValue={values?.user}
                />
                {isKey ? (
                    <ProFormText name="privatekey" label="密钥路径" width="md" rules={[{ required: true }]} initialValue={values?.privatekey} />
                ) : null}
                {isKey ? (
                    <ProFormText.Password
                        label="密钥密码"
                        name="key_passphrase"
                        width="md"
                        placeholder="输入则修改"
                    />
                ) : (
                    <ProFormText.Password
                        label="服务器密码"
                        name="password"
                        width="md"
                        placeholder="输入则修改"
                    />
                )}
            </ProForm.Group>
        </ModalForm>
    );
};

export default UpdateForm;