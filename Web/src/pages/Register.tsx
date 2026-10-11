import { useState } from "react";
import { App, Button, Card, Form, Input, Typography } from "antd";
import { LockOutlined, MobileOutlined, UserOutlined } from "@ant-design/icons";
import { Link, useNavigate } from "react-router-dom";
import styled from "@emotion/styled";
import { useAuth } from "../context/auth-context";

interface RegisterFormValues {
  phone: string;
  username: string;
  password: string;
  confirm_password: string;
}

const Wrapper = styled.div`
  min-height: calc(100vh - 64px);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
`;

const StyledCard = styled(Card)`
  width: 100%;
  max-width: 420px;
  background: rgba(255, 255, 255, 0.04) !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  backdrop-filter: blur(12px);

  .ant-card-body {
    padding: 32px;
  }
`;

const Title = styled(Typography.Title)`
  text-align: center;
  margin-bottom: 8px !important;
  color: #e5e7eb !important;
`;

const SubTitle = styled(Typography.Text)`
  display: block;
  text-align: center;
  margin-bottom: 24px;
  color: #9ca3af;
`;

export default function Register() {
  const { message } = App.useApp();
  const { register } = useAuth();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  const onFinish = async (values: RegisterFormValues) => {
    setLoading(true);
    try {
      await register(values);
      message.success("注册成功，请登录");
      navigate("/login");
    } catch (error) {
      message.error(error instanceof Error ? error.message : "注册失败");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Wrapper>
      <StyledCard>
        <Title level={3}>创建账号</Title>
        <SubTitle>加入 WindLiberity</SubTitle>
        <Form name="register" size="large" onFinish={onFinish} autoComplete="off">
          <Form.Item
            name="phone"
            rules={[
              { required: true, message: "请输入手机号" },
              { pattern: /^1[3-9]\d{9}$/, message: "手机号格式不正确" },
            ]}
          >
            <Input prefix={<MobileOutlined />} placeholder="手机号" />
          </Form.Item>
          <Form.Item
            name="username"
            rules={[{ required: true, min: 4, max: 18, message: "请输入 4-18 位用户名" }]}
          >
            <Input prefix={<UserOutlined />} placeholder="用户名" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, min: 6, max: 20, message: "请输入 6-20 位密码" }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="密码" />
          </Form.Item>
          <Form.Item
            name="confirm_password"
            dependencies={["password"]}
            rules={[
              { required: true, message: "请再次输入密码" },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue("password") === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("两次输入的密码不一致"));
                },
              }),
            ]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="确认密码" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block loading={loading}>
              注册
            </Button>
          </Form.Item>
        </Form>
        <div style={{ textAlign: "center" }}>
          <Typography.Text style={{ color: "#9ca3af" }}>
            已有账号？<Link to="/login">立即登录</Link>
          </Typography.Text>
        </div>
      </StyledCard>
    </Wrapper>
  );
}
