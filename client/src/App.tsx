import React from 'react';

const App: React.FC = () => {
  const get = async () => {
    const builds = await window.backend.getBuilds();
    console.log(builds);
  };

  return (
    <div id="app" className="App">
      <header className="App-header">
        <p>
          Welcome to your new <code>wails/react</code> project.
        </p>
        <button onClick={get}>click me</button>
      </header>
    </div>
  );
};

export default App;
