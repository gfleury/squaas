import React from 'react';

import { HashRouter as Router, Route, Link } from "react-router-dom";
import Grid from '@material-ui/core/Grid';

import DBqueryBench from 'd_bquery_bench';

import QueryList from "./QueryList";
import QueryNew from "./QueryNew";
import QueryEdit from "./QueryEdit"

import './App.css';

const isLocalhost = Boolean(
  window.location.hostname === 'localhost' ||
  // [::1] is the IPv6 localhost address.
  window.location.hostname === '[::1]' ||
  // 127.0.0.1/8 is considered localhost for IPv4.
  window.location.hostname.match(
    /^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/
  )
);

function App() {
  var api = new DBqueryBench.QueryApi();
  var dbApi = new DBqueryBench.DatabasesApi();

  api.apiClient.basePath = "/v1";
  dbApi.apiClient.basePath = "/v1";

  if (isLocalhost) {
    api.apiClient.basePath = "http://localhost:8080/v1";
    dbApi.apiClient.basePath = "http://localhost:8080/v1";
    api.apiClient.defaultHeaders = { "Authorization": "Basic YWRtaW46YWRtaW4=" }
    dbApi.apiClient.defaultHeaders = { "Authorization": "Basic YWRtaW46YWRtaW4=" }
  }

  return (
    <div className="App">
      <Router basename="/">
        <Header />
        <Route exact path="/" component={Home} />
        <Route exact path="/queries"
          render={(props) => <QueryList api={api} {...props} />}
        />
        <Route path="/queries/new"
          render={(props) => <QueryNew dbApi={dbApi} api={api} {...props} />}
        />
        <Route path="/queries/edit/:id"
          render={(props) => <QueryEdit dbApi={dbApi} api={api} {...props} />}
        />
      </Router>
    </div >
  );
}

function Home() {
  return <h2>Home</h2>;
}

function Header() {
  return (
    <Grid container spacing={3}>
      <Grid item xs>
        <Link to="/">Home</Link>
      </Grid>
      <Grid item xs>
        <Link to="/queries">Queries</Link>
      </Grid>
      <Grid item xs>
        <Link to="/queries/new">Create Query</Link>
      </Grid>
    </Grid>
  );
}

export default App;
