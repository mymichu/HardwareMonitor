import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Home from './component/PageHome';
import Cpu from './component/PageCPU';

function App() {

  return (
    <div className="App" style={{width:'100%'}}>
      <Router>
        <div>
          <Switch>
              <Route exact path='/' component={Home} />
              <Route path='/cpu' component={Cpu} />
          </Switch>
        </div>
      </Router>

    </div>
  );
}

export default App;
