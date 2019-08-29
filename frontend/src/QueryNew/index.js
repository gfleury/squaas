import React from 'react';
import { Redirect } from 'react-router-dom'

import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
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

export default class QueryNew extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            databases: [],
            parseStatus: false,
            includedId: 0,
            parseText: "",
            server: "",
            status: "Ready",
            ticketid: "",
            query: "-- SQL Query\n-- Paste it here\nSELECT * FROM XTABLE;",
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleParse = this.handleParse.bind(this);
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
        close: {
            padding: theme.spacing(0.5),
        },
    }));

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
    }

    handleChange(event) {
        if (event.target.name === "ticketid") {
            event.target.value = event.target.value.toUpperCase();
        }
        let change = {};
        change[event.target.name] = event.target.value;
        //console.log(change);
        this.setState(change);
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

    handleSubmit(event) {
        var _this = this;
        var query = new DBqueryBench.Query();

        query.status = this.state.status;
        query.query = this.state.query;
        query.ticketid = this.state.ticketid;
        query.servername = this.state.server;

        this.props.api.addQuery(query, function (error, data, response) {
            if (error) {
                _this.setState({ parseStatus: true, parseText: response.body.error })
            } else {
                _this.setState({ parseStatus: true, parseText: "Successfully included query, redirecting..." })
                _this.setState({ includedId: response.body.id });
            }
        });
        //console.log(this.state);
        event.preventDefault();
    }

    render() {
        if (this.state.includedId !== 0) {
            return <Redirect to={`/queries/edit/${this.state.includedId}`} />
        }
        return (
            <form className={this.classes.container} onSubmit={this.handleSubmit} autoComplete="off" >
                <Grid container spacing={3}>
                    <Grid item xs={3}>
                        <TextField
                            required
                            id="ticketid"
                            name="ticketid"
                            label="Ticket ID"
                            className={this.classes.textField}
                            value={this.state.ticketid}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                        />
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl className={this.classes.formControl}>
                            <InputLabel htmlFor="status">Query State</InputLabel>
                            <Select
                                required
                                value={this.state.status}
                                onChange={this.handleChange}
                                input={<Input name="status" id="status" />}
                            >
                                <MenuItem value="OnHold">On Hold</MenuItem>
                                <MenuItem value="Ready">Ready</MenuItem>
                            </Select>
                            <FormHelperText>Select the query state</FormHelperText>
                        </FormControl>
                    </Grid>
                    <Grid item xs={3} >
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
                    <Grid item xs={12}>
                        <TextField
                            required
                            id="query"
                            name="query"
                            label="SQL Query"
                            multiline
                            value={this.state.query}
                            className={this.classes.textField}
                            margin="normal"
                            variant="outlined"
                            InputLabelProps={{
                                shrink: true,
                            }}
                            onChange={this.handleChange}
                            fullWidth
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <SyntaxHighlighter
                            id="query"
                            name="query"
                            label="SQL Query"
                            showLineNumbers
                            onChange={this.handleChange}
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
                    </Grid>
                </Grid>
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
            </form>
        );
    }
}
