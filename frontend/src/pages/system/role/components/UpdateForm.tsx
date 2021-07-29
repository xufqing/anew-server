import React from 'react';
import { updateRole } from '@/services/anew/role';
import ProForm, { ModalForm, ProFormText } from '@ant-design/pro-form';
import { message } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

export type UpdateFormProps = {
    modalVisible: boolean;
    onChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.RoleList| undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, onChange, values } = props;

    return (
        <ModalForm
            title="更新角色"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                updateRole(v as any, values?.id).then((res) => {
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
                <ProFormText name="keyword" label="关键字" width="md" initialValue={values?.keyword} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText name="desc" label="说明" width="md" initialValue={values?.desc} />
            </ProForm.Group>
        </ModalForm>
    );
};

export default UpdateForm;