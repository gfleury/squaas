import React from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';

import EditIcon from '@material-ui/icons/Edit';

import DBqueryBench from 'd_bquery_bench';

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

function createData() {
    var query = new DBqueryBench.Query();
    query.ticketid = "INFRA-4940";
    query.id = 10;
    query.owner = "george.fleury@trustyou.com";
    query.query = "SELECT * FROM XTABLE;"
    query.status = "RUNNING";
    return query;
}

const rows = [
    createData(),
    createData(),
    createData(),
    createData(),
    createData(),
];

export default function SimpleTable({ match }) {
    const classes = useStyles();

    return (
        <Grid item xs={12}>
            <Paper className={classes.root}>
                <Table className={classes.table}>
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
                        {rows.map(row => (
                            <TableRow hover
                                key={row.name}>
                                <TableCell component="th" align="left">
                                    <Link to={`${match.url}/${row.id}`}>
                                        <EditIcon className={classes.rightIcon} />
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
