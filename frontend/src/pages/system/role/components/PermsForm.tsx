import React, { useEffect, useState } from 'react';
import { queryMenus } from '@/services/anew/menu';
import { queryApis } from '@/services/anew/api';
import { getRolePermsByID, updatePermsRole } from '@/services/anew/role';
import { DrawerForm } from '@ant-design/pro-form';
import { message, Tree, Checkbox, Col, Row, Divider } from 'antd';
import type { ActionType } from '@ant-design/pro-table';
import type { CheckboxValueType } from 'antd/lib/checkbox/Group'

const loopTreeItem = (tree: API.MenuList[]): API.MenuList[] =>
    tree.map(({ children, ...item }) => ({
        ...item,
        title: item.name,
        key: item.id,
        children: children && loopTreeItem(children),
    }));


export type PermsFormProps = {
    modalVisible: boolean;
    onChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.RoleList | undefined;
};

const PermsForm: React.FC<PermsFormProps> = (props) => {
    const { actionRef, modalVisible, onChange, values } = props;
    const [menuData, setMenuData] = useState<API.MenuList[]>([]);
    const [apiData, setApiData] = useState<API.ApiList[]>([]);
    const [checkedMenu, setCheckedMenu] = useState<React.Key[]>([]);
    const [checkedApi, setCheckedApi] = useState<CheckboxValueType[]>([]);

    const onCheck = (keys: any, info: any) => {
        let allKeys = keys.checked;
        const parentKey = info.node.parent_id;
        if (allKeys.indexOf(parentKey)) {
            setCheckedMenu(allKeys);
        } else {
            allKeys = allKeys.push(parentKey);
            setCheckedMenu(allKeys);
        }
    };

    const onCheckChange = (checkedValue: CheckboxValueType[]) => {
        //console.log(a.filter(function(v){ return !(b.indexOf(v) > -1) }).concat(b.filter(function(v){ return !(a.indexOf(v) > -1)})))
        setCheckedApi(checkedValue);
    };

    useEffect(() => {
        queryMenus().then((res) => setMenuData(loopTreeItem(res.data)));
        queryApis().then((res) => setApiData(res.data));
        getRolePermsByID(values?.id).then((res) => {
            setCheckedMenu(res.data.menus_id);
            setCheckedApi(res.data.apis_id);
        });
    }, []);
    return (
        <DrawerForm
            //title="设置权限"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async () => {
                updatePermsRole({
                    menus_id: checkedMenu,
                    apis_id: checkedApi,
                }, values?.id).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        if (actionRef.current) {
                            actionRef.current.reload();
                        }
                    }
                })
                return true;
            }}
        >
            <h3>菜单权限</h3>
            <Divider />

            <Tree
                checkable
                checkStrictly
                style={{ width: 330 }}
                //defaultCheckedKeys={selectedKeys}
                //defaultSelectedKeys={selectedKeys}
                autoExpandParent={true}
                selectable={false}
                onCheck={onCheck}
                checkedKeys={checkedMenu}
                treeData={menuData as any}
            />

            <Divider />
            <h3>API权限</h3>
            <Divider />
            <Checkbox.Group style={{ width: '100%' }} value={checkedApi} onChange={onCheckChange}>
                {apiData.map((item, index) => {
                    return (
                        <div key={index}>
                            <h4>{item.name}</h4>
                            <Row>
                                {item.children?.map((item, index) => {
                                    return (
                                        <Col span={4} key={index}>
                                            <Checkbox value={item.id}>{item.name}</Checkbox>
                                        </Col>
                                    );
                                })}
                            </Row>
                            <Divider />
                        </div>
                    );
                })}
            </Checkbox.Group>
        </DrawerForm>
    );
};

export default PermsForm;