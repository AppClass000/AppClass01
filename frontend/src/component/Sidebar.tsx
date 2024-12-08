import * as React from 'react'
import { SidebarData } from './SidebarData';

type Props = {
};

const Sidebar: React.FC<Props> = (props) => {
    return(
        <div className="Side-bar">
            <ul className='Sidebar-list'>
                {SidebarData.map((value, key) => {
                    return(
                        <li key={key}  className='row' onClick={() => {
                            window.location.pathname = value.link;
                        }}>
                            <div id='icon'>{value.icon}</div>
                            <div id='title'>{value.title}</div>
                        </li>
                    )
                })}
            </ul>
            {/* <p>現在{credit}単位</p> */}
        </div>
    );
};

export default Sidebar;