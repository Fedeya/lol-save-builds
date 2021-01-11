import React, { useEffect } from 'react';
import { MemoryRouter, Switch, Route } from 'react-router-dom';
import './output.css';

//  Components
import Layout from './components/Layout';

// Store
import { useBuildsStore } from './store/builds';

// Views
import Home from './views/Home';
import Build from './views/Build';

const App: React.FC = () => {
  const loadBuilds = useBuildsStore(state => state.loadBuilds);

  useEffect(() => {
    loadBuilds();
  }, [loadBuilds]);

  return (
    <MemoryRouter>
      <Switch>
        <Layout>
          <Route exact path="/" component={Home} />
          <Route path="/build/:id" component={Build} />
        </Layout>
      </Switch>
    </MemoryRouter>
  );
};

export default App;
