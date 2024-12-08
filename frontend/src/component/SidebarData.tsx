import React from 'react'
import EditCalendarIcon from '@mui/icons-material/EditCalendar';
import DnsIcon from '@mui/icons-material/Dns';
import HomeIcon from '@mui/icons-material/Home';

export const  SidebarData = [
    {
        title:"ホーム",
        icon: <HomeIcon />,
        link: "/home",
    },
    {
        title:"時間割",
        icon: <EditCalendarIcon />,
        link: "/timetable",
    },
    {
        title:"授業リスト",
        icon: <DnsIcon />,
        link: "/classes",
    },
    
];