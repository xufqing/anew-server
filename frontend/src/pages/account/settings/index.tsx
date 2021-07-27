import { GridContent } from '@ant-design/pro-layout';
import { UserOutlined, EditOutlined, UploadOutlined, MailOutlined, MobileOutlined, UsergroupAddOutlined, TeamOutlined } from '@ant-design/icons';
import React, { useState, useEffect } from 'react';
import { useModel } from 'umi';
import { Card, Col, Button, Upload, Row, Tabs, message } from 'antd';
import BaseForm from './components/BaseForm';
import type { RcFile, UploadChangeParam } from 'antd/lib/upload';
//import ChangePasswordFrom from './components/ChangePasswordFrom';
//import IconFont from '@/components/IconFont';
//import { querycurrentUser } from './service';
import styles from './settings.less';

const Settings: React.FC = () => {
    const { initialState } = useModel('@@initialState');

    if (!initialState || !initialState.currentUser) {
        return null;
    }

    const { currentUser } = initialState;
    const [uploadLoading, setUploadLoading] = useState(false);
    //const [currentUser, setcurrentUser] = useState({});

    //   useEffect(() => {
    //     getInfo();
    //   }, []);


    //   const getInfo = () => {
    //     querycurrentUser().then((res) => {
    //       setcurrentUser(res.data);
    //     });
    //   };

    const beforeUpload = (file: RcFile) => {
        const isJpgOrPng =
            file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/gif';
        if (!isJpgOrPng) {
            message.error('只可以上传JPG/PNG/GIF图片!');
        }
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isLt2M) {
            message.error('图片必须小于2MB!');
        }
        return isJpgOrPng && isLt2M;
    };

    const handleChange = (info: UploadChangeParam) => {
        if (info.file.status === 'uploading') {
            setUploadLoading(true);
            return;
        }
        if (info.file.status === 'done') {
            message.success('上传成功');
            // let currentUser = JSON.parse(localStorage.getItem('user')) || {};
            // currentUser.avatar = info.file.response.data.url;
            // localStorage.setItem('user', JSON.stringify(currentUser));
            //getInfo();
        }
    };

    const tokenHeaders = {
        Authorization: 'Bearer ' + localStorage.getItem('token'),
    };
    return (
        <GridContent>
            {currentUser.username && (
                <Row gutter={24}>
                    <Col span={10}>
                        <Card
                            title="关于我"
                            bordered={false}
                            style={{
                                marginBottom: 14,
                            }}
                        >
                            <div>
                                <div className={styles.avatarHolder}>
                                    <img alt="" src={currentUser.avatar} />
                                    <Upload
                                        name="avatar"
                                        headers={tokenHeaders}
                                        accept=".jpg,.jpeg,.png,.gif"
                                        className="avatar-uploader"
                                        showUploadList={false}
                                        action="/api/v1/user/info/uploadImg"
                                        beforeUpload={beforeUpload}
                                        onChange={handleChange}
                                    >
                                        <div className={styles.button_view}>
                                            <Button>
                                                <UploadOutlined /> 更换头像
                                            </Button>
                                        </div>
                                    </Upload>
                                </div>
                                <div className={styles.detail}>
                                    <div>
                                        <p style={{ marginRight: '15px' }}>
                                            <UserOutlined />
                                            用户名
                                        </p>
                                        {currentUser.username}
                                    </div>
                                    <div>
                                        <p style={{ marginRight: '29px' }}>
                                            <EditOutlined />
                                            姓名
                                        </p>
                                        {currentUser.name}
                                    </div>
                                    <div>
                                        <p style={{ marginRight: '29px' }}>
                                            <MailOutlined />
                                            邮箱
                                        </p>
                                        {currentUser.email}
                                    </div>
                                    <div>
                                        <p style={{ marginRight: '29px' }}>
                                            <MobileOutlined />
                                            手机
                                        </p>
                                        {currentUser.mobile}
                                    </div>
                                    <div>
                                        <p style={{ marginRight: '29px' }}>
                                            <UsergroupAddOutlined />
                                            角色
                                        </p>
                                        {currentUser.role?.name}
                                    </div>
                                    <div>
                                        <p style={{ marginRight: '29px' }}>
                                            <TeamOutlined />
                                            部门
                                        </p>
                                        {currentUser.dept?.name}
                                    </div>
                                </div>
                            </div>
                        </Card>
                    </Col>
                    <Col span={14}>
                        <Card title="个人设置" bordered={false} >
                            <Tabs tabPosition="right" onChange={() => { }}>
                                <Tabs.TabPane tab="基本信息" key="baseInfo">
                                    <BaseForm values={currentUser} />
                                </Tabs.TabPane>
                                <Tabs.TabPane tab="修改密码" key="changePwd">
                                    {/* <ChangePasswordFrom /> */}
                                </Tabs.TabPane>
                            </Tabs>
                        </Card>
                    </Col>
                </Row>
            )}
        </GridContent>
    );
};

export default Settings;
