import React from 'react';

import Sidebar from './Sidebar';
import Input from './Input';

const Layout: React.FC = ({ children }) => {
  return (
    <div className="flex min-h-screen">
      <Sidebar />
      <main className="w-full h-screen">
        <Input />
        <div className="w-full max-h-full overflow-auto">{children}</div>
      </main>
    </div>
  );
};

export default Layout;
