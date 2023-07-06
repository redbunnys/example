'use client'
import { FC } from "react";
import { Layout } from "@arco-design/web-react/es/";
import { Menu } from '@arco-design/web-react';
import { IconApps, IconBug, IconBulb } from '@arco-design/web-react/icon';
import styles from './layout.module.scss'
const { Header, Sider, Content, Footer } = Layout


interface DashBoardLayoutProps {
    children: React.ReactNode
}


const DashBoardLayouts: FC<DashBoardLayoutProps> = ({ children }) => {
    return (
        <>
            <Layout className={styles["layout-basic-demo"]}>
                <Header className={styles["arco-layout-header"]}>Header</Header>
                <Layout>
                    <Sider style={{"width":""}} className={styles["arco-layout-sider"]}><Menus/></Sider>
                    <Content className={styles["arco-layout-content"]}>{children}</Content>
                </Layout>
                <Footer className={styles["arco-layout-footer"]}>Footer</Footer>
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




