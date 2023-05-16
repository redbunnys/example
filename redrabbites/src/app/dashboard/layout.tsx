import { FC } from "react";
import DashBoardLayouts from '../components/dashboard/layout'
import '@arco-design/web-react/dist/css/arco.css';


interface DashBoardLayoutProps {
    children: React.ReactNode
}

export const metadata = {
    title: 'Dashboard',
    description: 'Dashboard app',
}

const DashBoardLayout: FC<DashBoardLayoutProps> = ({ children }) => {
    return (
        <>
            <DashBoardLayouts children={children}/>
        </>
    );
}

export default DashBoardLayout;