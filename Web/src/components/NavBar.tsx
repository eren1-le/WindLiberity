/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-18 00:47:55
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-18 20:11:38
 * @FilePath: /app/WindLiberity/Wind-Liberity/src/components/NavBar.tsx
 * @Description: 
 * 
 */
import React from "react";
import { Layout, Menu } from 'antd';
import type { MenuProps } from "antd";
import { Link, useLocation } from 'react-router-dom';
import styled from '@emotion/styled';
import { useState } from 'react';
import { HomeOutlined, DollarOutlined, EuroCircleFilled } from "@ant-design/icons";
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

    const items: MenuItem[] = [
    {
        label: '首页',
        key: 'home',
        icon: <HomeOutlined />,
    },
    {
        label: '番剧',
        key: 'anime',
        icon: <DollarOutlined />,
    }, 
    {
        label: '登陆/注册',
        key: 'login/register',
        icon: <EuroCircleFilled />,
    },
];
    const [current, setCurrent] = useState('home');     
    const onClick: MenuProps['onClick'] = (e) => {
        console.log('click', e);
        setCurrent(e.key);
    };

    return (
        <StyleHeader>
            <Logo>WindLiberity</Logo>
            <MainMenu mode='horizontal'
            items={items} 
            onClick={onClick}>
            selectedKeys={[current]}
            </MainMenu>
        </StyleHeader>
    );
}; 

export default NavBar;


     





