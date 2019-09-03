import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormLabel from '@material-ui/core/FormLabel';
import TextField from '@material-ui/core/TextField';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Input from '@material-ui/core/Input';

import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';

import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomDark } from 'react-syntax-highlighter/dist/esm/styles/prism';

import DBqueryBench from 'd_bquery_bench';


export default class QueryEdit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            databases: [],
            parseStatus: false,
            parseText: "",
            id: props.match.params.id,
            server: "",
            status: "ready",
            ticketid: "",
            query: "",
            owner: "",
            hasselect: false,
            hasalter: false,
            hastransaction: false,
            hasinsert: false,
            hasdelete: false,
            hasupdate: false,
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleParse = this.handleParse.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleApprove = this.handleApprove.bind(this);
        this.handleDisapprove = this.handleDisapprove.bind(this);
    }

    classes = makeStyles(theme => ({
        container: {
            display: 'flex',
            flexWrap: 'wrap',
        },
        textField: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
        },
        dense: {
            marginTop: theme.spacing(2),
        },
        menu: {
            width: 200,
        },
        root: {
            display: 'flex',
            flexWrap: 'wrap',
        },
        formControl: {
            margin: theme.spacing(1),
            minWidth: 120,
        },
        selectEmpty: {
            marginTop: theme.spacing(2),
        },
    }));

    handleChange(event) {
        if (event.target.name === "ticketid") {
            event.target.value = event.target.value.toUpperCase();
        }
        let change = {};
        change[event.target.name] = event.target.value;
        //console.log(change);
        this.setState(change);
    }

    componentDidMount() {
        var _this = this;
        this.props.dbApi.getDatabases(function (error, data) {
            if (error) {
                console.error(error);
            } else {
                console.log('API called successfully.');
                console.log(data);
                _this.setState({ databases: data })
            }
        });

        this.props.api.getQueryById(this.state.id, function (error, data) {
            if (error) {
                console.error(error);
            } else {
                console.log('API called successfully.');
                console.log(data);
                _this.setState({
                    server: data.servername,
                    status: data.status,
                    ticketid: data.ticketid,
                    query: data.query,
                    owner: data.owner.name,
                    hasselect: data.hasselect,
                    hasalter: data.hasalter,
                    hastransaction: data.hastransaction,
                    hasinsert: data.hasinsert,
                    hasdelete: data.hasdelete,
                    hasupdate: data.hasupdate,
                })
            }
        });

    }

    handleParse(event) {
        var _this = this;
        var query = new DBqueryBench.Query();

        query.status = "PARSEONLY"
        query.query = this.state.query;
        query.ticketid = this.state.ticketid;
        query.servername = this.state.server;

        this.props.api.addQuery(query, function (error, data, response) {
            if (error) {
                _this.setState({ parseStatus: true, parseText: response.body.error })
            } else {
                _this.setState({ parseStatus: true, parseText: "Successfully checked" })
                _this.setState({})
            }
        });
        //console.log(this.state);
        event.preventDefault();
    }

    handleApprove(event) {
        var _this = this;

        this.props.api.approveQuery(this.state.id, "pending", function (error, data, response) {
            if (error) {
                _this.setState({ parseStatus: true, parseText: response.body.error })
            } else {
                _this.setState({ parseStatus: true, parseText: "Successfully approved" })
                _this.setState({})
            }
        });
        //console.log(this.state);
        event.preventDefault();
    }

    handleDisapprove(event) {
        var _this = this;

        this.props.api.deleteApprovalQuery(this.state.id, function (error, data, response) {
            if (error) {
                _this.setState({ parseStatus: true, parseText: response.body.error })
            } else {
                _this.setState({ parseStatus: true, parseText: "Successfully disapproved" })
                _this.setState({})
            }
        });
        //console.log(this.state);
        event.preventDefault();
    }

    handleSubmit(event) {
        var _this = this;
        var query = new DBqueryBench.Query();

        query.status = this.state.status;
        query.query = this.state.query;
        query.ticketid = this.state.ticketid;
        query.servername = this.state.server;
        query.id = this.state.id

        this.props.api.updateQuery(query, function (error, data, response) {
            if (error) {
                _this.setState({ parseStatus: true, parseText: response.body.error })
            } else {
                _this.setState({ parseStatus: true, parseText: "Successfully saved query" })
                _this.componentDidMount()
            }
        });
        //console.log(this.state);
        event.preventDefault();
    }

    render() {
        return (

            <form className={this.classes.container} onSubmit={this.handleSubmit} autoComplete="off" >
                <Grid container spacing={3}>
                    <Grid item xs={6} sm={3}>
                        <TextField
                            disabled
                            id="outlined-name"
                            label="ID"
                            name="id"
                            className={this.classes.textField}
                            value={this.state.id}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                        />
                    </Grid>
                    <Grid item xs={6} sm={3}>
                        <FormControl className={this.classes.formControl}>
                            <InputLabel htmlFor="status">Query State</InputLabel>
                            <Select
                                required
                                value={this.state.status}
                                onChange={this.handleChange}
                                input={<Input name="status" id="status" />}
                            >
                                <MenuItem value="pending">On Hold</MenuItem>
                                <MenuItem value="ready">Ready</MenuItem>
                            </Select>
                            <FormHelperText>Select the query state</FormHelperText>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6} sm={3}>
                        <TextField
                            disabled
                            id="outlined-name"
                            label="Owner"
                            name="owner"
                            className={this.classes.textField}
                            value={this.state.owner}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                        />
                    </Grid>
                    <Grid item xs={6} sm={3}>
                        <TextField
                            required
                            id="outlined-name"
                            label="Ticket ID"
                            name="ticketid"
                            className={this.classes.textField}
                            value={this.state.ticketid}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                        />
                    </Grid>
                    <Grid item xs={9}>

                        <FormLabel component="legend">Query behaviors</FormLabel>
                        <FormControlLabel
                            value="Transaction"
                            control={<Checkbox checked={this.state.hastransaction} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="Transaction"
                            labelPlacement="top"
                        />
                        <FormControlLabel
                            value="INSERT"
                            control={<Checkbox checked={this.state.hasinsert} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="INSERT"
                            labelPlacement="top"
                        />
                        <FormControlLabel
                            value="UPDATE"
                            control={<Checkbox checked={this.state.hasupdate} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="UPDATE"
                            labelPlacement="top"
                        />
                        <FormControlLabel
                            value="DELETE"
                            control={<Checkbox checked={this.state.hasdelete} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="DELETE"
                            labelPlacement="top"
                        />
                        <FormControlLabel
                            value="SELECT"
                            control={<Checkbox checked={this.state.hasselect} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="SELECT"
                            labelPlacement="top"
                        />
                        <FormControlLabel
                            value="ALTER"
                            control={<Checkbox checked={this.state.hasalter} inputProps={{
                                'aria-label': 'disabled checked checkbox',
                            }} disabled color="primary" />}
                            label="ALTER"
                            labelPlacement="top"
                        />

                    </Grid>
                    <Grid item xs={3}>
                        <FormControl className={this.classes.formControl}>
                            <InputLabel htmlFor="server">Database Server</InputLabel>
                            <Select
                                required
                                value={this.state.server}
                                onChange={this.handleChange}
                                input={<Input name="server" id="server" />}
                            >
                                <MenuItem value="">
                                    <em>None</em>
                                </MenuItem>
                                {this.state.databases.map(db => (
                                    <MenuItem value={db.name} key={db.name}>{db.name}</MenuItem>
                                ))}
                            </Select>
                            <FormHelperText>Select the database for running the query</FormHelperText>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <TextField
                            required
                            id="query"
                            name="query"
                            label="SQL Query"
                            multiline
                            value={this.state.query}
                            className={this.classes.textField}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                            InputLabelProps={{
                                shrink: true,
                            }}
                            fullWidth
                        />
                    </Grid>
                    <Grid item xs={6}>
                        <SyntaxHighlighter
                            id="query"
                            name="query"
                            label="SQL Query"
                            showLineNumbers
                            language="sql"
                            style={atomDark}
                            children={this.state.query}
                        />
                    </Grid>
                    <Grid item xs>
                        <Button onClick={this.handleParse} variant="contained" className={this.classes.button}>
                            PARSE
                </Button>
                        <Button type="submit" variant="contained" className={this.classes.button}>
                            SAVE
                </Button>
                        <Button onClick={this.handleApprove} variant="contained" className={this.classes.button}>
                            APPROVE
                </Button>
                        <Button onClick={this.handleDisapprove} variant="contained" className={this.classes.button}>
                            DISAPROVE
                </Button>
                    </Grid>
                </Grid >
                <Snackbar
                    anchorOrigin={{
                        vertical: 'bottom',
                        horizontal: 'left',
                    }}
                    open={this.state.parseStatus}
                    autoHideDuration={2000}
                    onClose={() => { this.setState({ parseStatus: false }) }}
                    ContentProps={{
                        'aria-describedby': 'message-id',
                    }}
                    message={<span id="message-id">{this.state.parseText}</span>}
                    action={[
                        <IconButton
                            key="close"
                            aria-label="close"
                            color="inherit"
                            className={this.classes.close}
                            onClick={() => { this.setState({ parseStatus: false }) }}
                        >
                            <CloseIcon />
                        </IconButton>,
                    ]}
                />
            </form >

        );
    }
}
