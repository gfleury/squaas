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
import Chip from '@material-ui/core/Chip';


import Moment from 'react-moment';

import MoreHorizIcon from '@material-ui/icons/MoreHoriz';

import TransformIcon from '@material-ui/icons/Transform';
import ChromeReaderModeIcon from '@material-ui/icons/ChromeReaderMode';
import OpenInBrowserIcon from '@material-ui/icons/OpenInBrowser';
import UpdateIcon from '@material-ui/icons/Update';
import DeleteForeverIcon from '@material-ui/icons/DeleteForever';
import MoodBadIcon from '@material-ui/icons/MoodBad';
import MoodIcon from '@material-ui/icons/Mood';





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

    getBehavior(query) {
        let queryBehaviors = (<></>);
        if (query.hastransaction) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="Transaction"
                size="small"
                icon={<TransformIcon />} /></>)
        }

        if (query.hasselect) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="Select"
                size="small"
                icon={<ChromeReaderModeIcon />} />
            </>)
        }

        if (query.hasinsert) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="Insert"
                size="small"
                icon={<OpenInBrowserIcon />} /></>)
        }

        if (query.hasupdate) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="Update"
                size="small"
                icon={<UpdateIcon />} /></>)
        }

        if (query.hasdelete) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="Delete"
                size="small"
                icon={<DeleteForeverIcon />} /></>)
        }

        if (query.hasalter) {
            queryBehaviors = (<>{queryBehaviors}<Chip
                label="DDL Modification"
                size="small"
                icon={<MoodBadIcon />} /></>)
        }

        return queryBehaviors
    }

    getApprovals(approvals) {
        let approved = 0
        let disaproved = 0
        console.log(approvals);
        approvals.map(approval => {
            if (approval.approved) {
                approved++;
            } else {
                disaproved++;
            }
            return true;
        })

        return (
            <>
                <Chip
                    label={approved}
                    size="small"
                    color="primary"
                    icon={<MoodIcon />} />
                <Chip
                    label={disaproved}
                    size="small"
                    color="secondary"
                    icon={<MoodBadIcon />} />
            </>
        )
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
                                <TableCell align="right">Target database</TableCell>
                                <TableCell align="right">Approval Count</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {
                                this.state.queries.map(query => (
                                    <TableRow hover
                                        key={query.id}>
                                        <TableCell component="th" align="left">
                                            <Link to={`${this.props.match.url}/edit/${query.id}`}>
                                                <MoreHorizIcon className={this.classes.rightIcon} />
                                            </Link>
                                        </TableCell>
                                        <TableCell component="th" scope="row" align="center">
                                            {query.ticketid}
                                        </TableCell>
                                        <TableCell align="center">{query.status}</TableCell>
                                        <TableCell align="right">{query.owner.name}</TableCell>
                                        <TableCell align="right"><Moment fromNow>{query.updatedAt}</Moment></TableCell>
                                        <TableCell align="right">{this.getBehavior(query)}</TableCell>
                                        <TableCell align="right">{query.servername}</TableCell>
                                        <TableCell align="right">{this.getApprovals(query.approvals)}</TableCell>
                                    </TableRow>
                                ))}
                        </TableBody>
                    </Table>
                </Paper>
            </Grid>
        );
    }
}