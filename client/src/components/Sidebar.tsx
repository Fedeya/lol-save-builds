import React from 'react';
import { NavLink as Link } from 'react-router-dom';

import { useBuildsStore } from '../store/builds';

const Sidebar: React.FC = () => {
  const builds = useBuildsStore(state => state.builds);

  return (
    <aside className="w-20 overflow-auto shadow-lg bg-gray-900 p-3 shadow-lg text-white min-h-screen flex flex-col items-center">
      <button className="bg-transparent focus:outline-none mb-3">
        <svg
          className="w-6 h-6"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
      </button>
      <button className="bg-transparent focus:outline-none">
        <svg
          className="w-6 h-6"
          data-darkreader-inline-fill=""
          data-darkreader-inline-stroke=""
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          />
        </svg>
      </button>

      <div className="flex flex-col items-center w-full h-full justify-center">
        {builds.map(build => (
          <Link
            key={build.id}
            to={`/build/${build.id}`}
            className="bg-transparent w-6 h-6 focus:outline-none focus:border-violet-300 rounded-full mb-3"
            activeClassName="focus:shadow-outline"
          >
            <img
              className="rounded-full w-6 h-6"
              src={build.image_url}
              alt={build.name}
            />
          </Link>
        ))}
      </div>
    </aside>
  );
};

export default Sidebar;
