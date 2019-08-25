import React from 'react';

import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Grid from '@material-ui/core/Grid';

import Edit from "./QueryEdit"

import List from "./QueryList"

import './App.css';

function App() {

  return (
    <div className="App">
      <Router basename="/frontend">
        <Header />
        <Route exact path="/" component={Home} />
        <Route path="/about" component={About} />
        <Route exact path="/queries" component={List} />
        <Route path="/queries/:id" component={Edit} />
      </Router>
    </div >
  );
}

function Home() {
  return <h2>Home</h2>;
}

function About() {
  return <h2>About</h2>;
}

function Header() {
  return (
    <Grid container spacing={3}>
      <Grid item xs>
        <Link to="/">Home</Link>
      </Grid>
      <Grid item xs>
        <Link to="/about">About</Link>
      </Grid>
      <Grid item xs>
        <Link to="/queries">Queries</Link>
      </Grid>
    </Grid>
  );
}

export default App;
