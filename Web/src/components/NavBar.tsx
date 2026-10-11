/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-18 00:47:55
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-18 20:11:38
 * @FilePath: /app/WindLiberity/Wind-Liberity/src/components/NavBar.tsx
 * @Description: 
 * 
 */
import { Layout, Menu, Dropdown, Avatar, Button, Space } from 'antd';
import type { MenuProps } from "antd";
import { Link, useLocation, useNavigate } from 'react-router-dom';
import styled from '@emotion/styled';
import { useEffect, useState } from 'react';
import { HomeOutlined, DollarOutlined, LoginOutlined, LogoutOutlined, UserOutlined } from "@ant-design/icons";
import { useAuth } from '../context/auth-context';
const { Header } = Layout;

const StyleHeader = styled(Header) `
    background: rgba(9, 9, 9, 0.25);
    backdrop-filter: blur(10px);
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    align-items: center;
    position: sticky;
    top: 0;
    z-index: 100;
`;

const StyledMenu = styled(Menu)`
  background: transparent !important;

  .ant-menu-item {
    color: #747a74 !important;
    
    &:hover {
      color: #585e59 !important;
    }
    
    .anticon {
      color: #717771 !important;
    }
    
    a {
      color: #737873 !important;
      &:hover {
        color: #e5eae6 !important;
      }
    }
  }
`;
const Logo = styled.div`
  color: #6e7d6e;
  font-weight: bold;
  font-size: 18px;
  margin-right: 24px;
`;

const  MainMenu = styled(StyledMenu)`
    flex: 1;
`;


const NavBar = () => {
    type MenuItem = Required<MenuProps>['items'][number];
    const { user, logout } = useAuth();
    const navigate = useNavigate();
    const location = useLocation();

    const items: MenuItem[] = [
        {
            label: <Link to="/">首页</Link>,
            key: 'home',
            icon: <HomeOutlined />,
        },
        {
            label: <Link to="/">番剧</Link>,
            key: 'anime',
            icon: <DollarOutlined />,
        },
    ];

    const userMenuItems: MenuProps['items'] = [
        {
            key: 'logout',
            label: '退出登录',
            icon: <LogoutOutlined />,
        },
    ];

    const [current, setCurrent] = useState('home');

    useEffect(() => {
        if (location.pathname.startsWith('/login') || location.pathname.startsWith('/register')) {
            setCurrent('login');
        } else {
            setCurrent('home');
        }
    }, [location.pathname]);

    const onClick: MenuProps['onClick'] = (e) => {
        setCurrent(e.key);
    };

    const onUserMenuClick: MenuProps['onClick'] = ({ key }) => {
        if (key === 'logout') {
            logout();
            navigate('/');
        }
    };

    return (
        <StyleHeader>
            <Logo>WindLiberity</Logo>
            <MainMenu mode='horizontal'
                items={items}
                onClick={onClick}
                selectedKeys={[current]}
            />
            {user ? (
                <Dropdown menu={{ items: userMenuItems, onClick: onUserMenuClick }}>
                    <Space style={{ cursor: 'pointer', color: '#e5e7eb' }}>
                        <Avatar size="small" icon={<UserOutlined />} style={{ background: '#fb7299' }} />
                        <span>{user.nickname || user.username}</span>
                    </Space>
                </Dropdown>
            ) : (
                <Link to="/login">
                    <Button type="primary" icon={<LoginOutlined />}>
                        登录 / 注册
                    </Button>
                </Link>
            )}
        </StyleHeader>
    );
}; 

export default NavBar;


     





