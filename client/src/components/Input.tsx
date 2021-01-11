import React, { useState } from 'react';
import { useBuildsStore } from '../store/builds';

const Input: React.FC = () => {
  const addBuild = useBuildsStore(state => state.addBuild);
  const [url, setUrl] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await addBuild(url);
  };

  return (
    <div className="w-full bg-gray-900 p-5">
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          className="focus:outline-none p-2 px-3 focus:border-teal-300 rounded-full focus:shadow-outline w-full text-white bg-gray-700"
          placeholder="URL to Save Build"
          value={url}
          onChange={e => setUrl(e.target.value)}
        />
      </form>
    </div>
  );
};

export default Input;
