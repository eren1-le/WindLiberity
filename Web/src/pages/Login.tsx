import { useState } from "react";
import { App, Button, Card, Form, Input, Typography } from "antd";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Link, useNavigate } from "react-router-dom";
import styled from "@emotion/styled";
import { useAuth } from "../context/auth-context";

interface LoginFormValues {
  username: string;
  password: string;
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

export default function Login() {
  const { message } = App.useApp();
  const { login } = useAuth();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  const onFinish = async (values: LoginFormValues) => {
    setLoading(true);
    try {
      await login(values);
      message.success("登录成功");
      navigate("/");
    } catch (error) {
      message.error(error instanceof Error ? error.message : "登录失败");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Wrapper>
      <StyledCard>
        <Title level={3}>欢迎回来</Title>
        <SubTitle>登录 WindLiberity</SubTitle>
        <Form name="login" size="large" onFinish={onFinish} autoComplete="off">
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
          <Form.Item>
            <Button type="primary" htmlType="submit" block loading={loading}>
              登录
            </Button>
          </Form.Item>
        </Form>
        <div style={{ textAlign: "center" }}>
          <Typography.Text style={{ color: "#9ca3af" }}>
            还没有账号？<Link to="/register">立即注册</Link>
          </Typography.Text>
        </div>
      </StyledCard>
    </Wrapper>
  );
}
