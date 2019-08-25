import React, { Component } from 'react';
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

const useStyles = makeStyles(theme => ({
    root: {
        width: '100%',
        marginTop: theme.spacing(3),
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
}));

export default class SimpleTable extends Component {
    constructor() {
        super();
        this.classes = useStyles();
        this.state = {
            rows: []
        }
    }

    componentDidMount() {
        const that = this;
        fetch('/v1/queris').then(result => {
            return result.json();
        }).then(data => {
            that.setState({rows: data.result});
        })
    }

    render() {
        if (!this.state.rows) {
            return '';
        }

        return (
            <Grid item xs={12}>
                <Paper className={this.classes.root}>
                    <Table className={this.classes.table}>
                        <TableHead>
                            <TableRow key="header">
                                <TableCell align="left"></TableCell>
                                <TableCell align="center">Ticket ID</TableCell>
                                <TableCell align="center">Status</TableCell>
                                <TableCell align="right">Owner</TableCell>
                                <TableCell align="right">Query behavior</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.rows.map(row => (
                                <TableRow hover
                                    key={row.name}>
                                    <TableCell component="th" align="left">
                                        <Link to={`${this.props.match.url}/${row.id}`}>
                                            <EditIcon className={this.classes.rightIcon} />
                                        </Link>
                                    </TableCell>
                                    <TableCell component="th" scope="row" align="center">
                                        {row.ticketid}
                                    </TableCell>
                                    <TableCell align="center">{row.status}</TableCell>
                                    <TableCell align="right">{row.owner}</TableCell>
                                    <TableCell align="right">{row.hastransaction}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </Paper>
            </Grid>
        );
    }
}
