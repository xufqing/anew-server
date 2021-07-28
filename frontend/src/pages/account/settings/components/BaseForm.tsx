import React from 'react';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { updateUserInfo } from '@/services/anew/user';
import { useModel } from 'umi';
import { message } from 'antd';


export type BaseFormProps = {
    values?: API.UserInfo;
};

const BaseForm: React.FC<BaseFormProps> = (props) => {
    const { values } = props;
    const { initialState, setInitialState } = useModel('@@initialState');

    return (
        <ProForm
            onFinish={async (v: API.UserInfo) => {
                await updateUserInfo(v, values?.id).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        let userInfo = {};
                        Object.assign(userInfo, initialState?.currentUser);
                        (userInfo as API.UserInfo).name = v.name;
                        (userInfo as API.UserInfo).mobile = v.mobile;
                        (userInfo as API.UserInfo).email = v.email
                        setInitialState({ ...initialState, currentUser: userInfo as API.UserInfo });
                    }
                });
            }}
        >
            <ProForm.Group>
                <ProFormText
                    name="name"
                    label="姓名"
                    width="md"
                    rules={[{ required: true }]}
                    initialValue={values?.name}
                />
                <ProFormText
                    name="mobile"
                    label="手机"
                    width="md"
                    rules={[
                        {
                            pattern: /^1(?:70\d|(?:9[89]|8[0-24-9]|7[135-8]|66|5[0-35-9])\d|3(?:4[0-8]|[0-35-9]\d))\d{7}$/,
                            message: '请输入正确的手机号码',
                        },
                    ]}
                    initialValue={values?.mobile}
                />
                <ProFormText
                    name="email"
                    label="邮箱"
                    width="md"
                    rules={[
                        {
                            type: 'email',
                            message: '请输入正确的邮箱地址',
                        },
                    ]}
                    initialValue={values?.email}
                />
            </ProForm.Group>
        </ProForm>
    );
};

export default BaseForm;
