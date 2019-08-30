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


import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomDark } from 'react-syntax-highlighter/dist/esm/styles/prism';

//import DBqueryBench from 'd_bquery_bench';


export default class QueryEdit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            databases: [],
            parseStatus: false,
            parseText: "",
            id: props.match.params.id,
            server: "",
            status: "Ready",
            ticketid: "",
            query: "-- SQL Query\n-- Paste it here\nSELECT * FROM XTABLE;",
            owner: "",
            hasselect: false,
            hasalter: false,
            hastransaction: false,
            hasinsert: false,
            hasdelete: false,
            hasupdate: false,
        };

        this.handleChange = this.handleChange.bind(this);
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

    render() {
        return (

            <form className={this.classes.container} noValidate autoComplete="off" >
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
                        <TextField
                            disabled
                            id="outlined-name"
                            label="Status"
                            name="status"
                            className={this.classes.textField}
                            value={this.state.status}
                            onChange={this.handleChange}
                            margin="normal"
                            variant="outlined"
                        />
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
                    <Grid item xs={12}>
                        <TextField
                            required
                            id="query"
                            name="query"
                            label="SQL Query"
                            multiline
                            defaultValue={this.state.query}
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
                    <Grid item xs={12}>
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
                        <Button type="submit" variant="contained" className={this.classes.button}>
                            PARSE
                </Button>
                        <Button type="submit" variant="contained" className={this.classes.button}>
                            SAVE
                </Button>
                        <Button type="submit" variant="contained" className={this.classes.button}>
                            APPROVE
                </Button>
                        <Button type="submit" variant="contained" className={this.classes.button}>
                            DISAPROVE
                </Button>
                    </Grid>
                </Grid >
            </form >

        );
    }
}
