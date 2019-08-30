import React from 'react';
/*eslint no-unused-vars: ["error", { "varsIgnorePattern": "Route" }]*/
import { BrowserRouter as Route, Router, Link } from "react-router-dom";

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';

import EditIcon from '@material-ui/icons/Edit';

import Moment from 'react-moment';

export default class QueryList extends React.Component {

    classes = makeStyles(theme => ({
        root: {
            width: '100%',
            marginTop: theme.spacing(3),
            overflowX: 'auto',
        },
        table: {
            minWidth: 650,
        },
    }));

    state = {
        queries: [],
    }

    componentDidMount() {
        var _this = this;
        this.props.api.getQueries({}, function (error, data) {
            if (error) {
                console.error(error);
            } else {
                console.log('API called successfully.');
                console.log(data);
                _this.setState({ queries: data })
            }
        });
    }

    render() {
        return (
            <Grid item xs={12} >
                <Paper className={this.classes.root}>
                    <Table className={this.classes.table}>
                        <TableHead>
                            <TableRow key="header">
                                <TableCell align="left"></TableCell>
                                <TableCell align="center">Ticket ID</TableCell>
                                <TableCell align="center">Status</TableCell>
                                <TableCell align="right">Owner</TableCell>
                                <TableCell align="right">Last update at</TableCell>
                                <TableCell align="right">Query behavior</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.queries.map(query => (
                                <TableRow hover
                                    key={query.id}>
                                    <TableCell component="th" align="left">
                                        <Link to={`${this.props.match.url}/edit/${query.id}`}>
                                            <EditIcon className={this.classes.rightIcon} />
                                        </Link>
                                    </TableCell>
                                    <TableCell component="th" scope="row" align="center">
                                        {query.ticketid}
                                    </TableCell>
                                    <TableCell align="center">{query.status}</TableCell>
                                    <TableCell align="right">{query.owner.name}</TableCell>
                                    <TableCell align="right"><Moment fromNow>{query.updatedAt}</Moment></TableCell>
                                    <TableCell align="right">{query.hastransaction}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
            </Grid>
        );
    }
}