'use client'
import { FC } from "react";
interface DashBoardLayoutProps {
    children: React.ReactNode
}

const DashBoardLayout: FC<DashBoardLayoutProps> = ({ children }) => {
    return (
        <html lang="en">
            <body >
                {children}
            </body>
        </html>
    );
}

export default DashBoardLayout;