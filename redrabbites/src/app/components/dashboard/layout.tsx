'use client'
import React, { FC } from "react";
import NextImg from 'next/image'
import { Avatar, Button, Dropdown, Layout, Image, Input, Space } from "@arco-design/web-react/es/";
import { Menu } from '@arco-design/web-react';
import { IconApps, IconBug, IconBulb, IconDown, IconPoweroff, IconSettings } from '@arco-design/web-react/icon';
import styles from './layout.module.scss'
const { Header, Sider, Content, Footer } = Layout
const InputSearch = Input.Search;

interface DashBoardLayoutProps {
  children: React.ReactNode
}


const DashBoardLayouts: FC<DashBoardLayoutProps> = ({ children }) => {
  return (
    <>
      <Layout className={styles["layout-basic-demo"]}>
        <Header className={styles["arco-layout-header"]}><Headers></Headers></Header>
        <Layout>
          <Sider style={{ "width": "" }} className={styles["arco-layout-sider"]}><Menus /></Sider>
          <Layout>
            <Content className={styles["arco-layout-content"]}>
              <div className={styles.childrenContent}>
              {children}
            </div>
            </Content>
            <Footer className={styles["arco-layout-footer"]}>redbunny </Footer>
          </Layout>
        </Layout>

      </Layout>

    </>
  );
}

export default DashBoardLayouts;


const MenuItem = Menu.Item;
const SubMenu = Menu.SubMenu;
const MenuItemGroup = Menu.ItemGroup;

const Menus = () => {
  return (
    <div className='menu-demo' style={{ height: "100%" }}>
      <Menu
        style={{ width: 250, height: '100%' }}
        hasCollapseButton
        defaultOpenKeys={['0']}
        defaultSelectedKeys={['0_1']}
      >
        <SubMenu
          key='0'
          title={
            <>
              <IconApps /> Navigation 1
            </>
          }
        >
          <MenuItem key='0_0'>Menu 1</MenuItem>
          <MenuItem key='0_1'>Menu 2</MenuItem>
          <MenuItem key='0_2' disabled>
            Menu 3
          </MenuItem>
        </SubMenu>
        <SubMenu
          key='1'
          title={
            <>
              <IconBug /> Navigation 2
            </>
          }
        >
          <MenuItem key='1_0'>Menu 1</MenuItem>
          <MenuItem key='1_1'>Menu 2</MenuItem>
          <MenuItem key='1_2'>Menu 3</MenuItem>
        </SubMenu>
        <SubMenu
          key='2'
          title={
            <>
              <IconBulb /> Navigation 3
            </>
          }
        >
          <MenuItemGroup key='2_0' title='Menu Group 1'>
            <MenuItem key='2_0_0'>Menu 1</MenuItem>
            <MenuItem key='2_0_1'>Menu 2</MenuItem>
          </MenuItemGroup>
          <MenuItemGroup key='2_1' title='Menu Group 1'>
            <MenuItem key='2_1_0'>Menu 3</MenuItem>
            <MenuItem key='2_1_1'>Menu 4</MenuItem>
          </MenuItemGroup>
        </SubMenu>
      </Menu>
    </div>
  );
};


<<<<<<< HEAD


=======
const dropList = (
  <Menu>
    <Menu.Item key='1'>
      <IconSettings />
      修改
    </Menu.Item>

    <Menu.Item key='2' style={{ "borderTop": "1px solid rgb(var(--gray-3))" }}>
      <IconPoweroff className={styles['dropdown-icon']} />注销
    </Menu.Item>
  </Menu>
);


function Headers() {
  return (
    <div className={styles.navbar}>
      <div className={styles.left}>
        <div>
          <Image
            width={40}
            src='//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a8c8cdb109cb051163646151a4a5083b.png~tplv-uwbnlip3yd-webp.webp'
            alt='lamp'
          />
        </div>

        <div>redbunny 后台</div>
      </div>
      <ul className={styles.right}>
        <li>
          <InputSearch
            searchButton
            defaultValue='Search content'
            placeholder='Enter keyword to search'
            style={{ width: 350 }}
          />
        </li>
        <li>主题</li>
        <li>中英文</li>
        <li>
          <Dropdown droplist={dropList} position='bl'>
            <Avatar>
              <img
                alt='avatar'
                src='//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp'
              />
            </Avatar>
          </Dropdown></li>
      </ul>
    </div>

  );
}



const ChildrenContent: React.FC<DashBoardLayoutProps> = ({ children }) => {
  return (
    <>

    </>
  );
}
>>>>>>> 2d57d9aa41c76cc3bbb26fe20a9496a07f814ea9
