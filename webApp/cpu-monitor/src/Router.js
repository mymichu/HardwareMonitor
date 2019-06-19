import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import { Route, BrowserRouter as Router } from 'react-router-dom'
import App from './App'
import CpuInfo from './component/PageCPU'

const routing = (
  <Router>
    <div>
      <Route exact path="/" component={App} />
      <Route path="/cpu" component={CpuInfo} />
    </div>
  </Router>
)

ReactDOM.render(routing, document.getElementById('root'))