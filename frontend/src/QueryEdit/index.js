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

const useStyles = makeStyles(theme => ({
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

export default function OutlinedTextFields({ match }) {
    const classes = useStyles();
    const [values, setValues] = React.useState({
        name: match.params.id,
        age: '',
        multiline: 'Controlled',
        currency: 'EUR',
    });

    const handleChange = name => event => {
        setValues({ ...values, [name]: event.target.value });
    };

    function xhandleChange(event) {
        setValues(oldValues => ({
            ...oldValues,
            [event.target.name]: event.target.value,
        }));
    }

    return (
        <form className={classes.container} noValidate autoComplete="off">
            <Grid item xs={6} sm={3}>
                <TextField
                    disabled
                    id="outlined-name"
                    label="ID"
                    className={classes.textField}
                    value={values.name}
                    onChange={handleChange('name')}
                    margin="normal"
                    variant="outlined"
                />
            </Grid>
            <Grid item xs={6} sm={3}>
                <TextField
                    disabled
                    id="outlined-name"
                    label="Status"
                    className={classes.textField}
                    value={values.name}
                    onChange={handleChange('name')}
                    margin="normal"
                    variant="outlined"
                />
            </Grid>
            <Grid item xs={6} sm={3}>
                <TextField
                    disabled
                    id="outlined-name"
                    label="Owner"
                    className={classes.textField}
                    value={values.name}
                    onChange={handleChange('name')}
                    margin="normal"
                    variant="outlined"
                />
            </Grid>
            <Grid item xs={6} sm={3}>
                <TextField
                    required
                    id="outlined-name"
                    label="Ticket ID"
                    className={classes.textField}
                    value={values.name}
                    onChange={handleChange('name')}
                    margin="normal"
                    variant="outlined"
                />
            </Grid>
            <Grid item xs={9}>
                <FormLabel component="legend">Query behaviors</FormLabel>
                <FormControlLabel
                    value="Transaction"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="Transaction"
                    labelPlacement="top"
                />
                <FormControlLabel
                    value="INSERT"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="INSERT"
                    labelPlacement="top"
                />
                <FormControlLabel
                    value="UPDATE"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="UPDATE"
                    labelPlacement="top"
                />
                <FormControlLabel
                    value="DELETE"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="DELETE"
                    labelPlacement="top"
                />
                <FormControlLabel
                    value="SELECT"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="SELECT"
                    labelPlacement="top"
                />
                <FormControlLabel
                    value="ALTER"
                    control={<Checkbox checked={true} inputProps={{
                        'aria-label': 'disabled checked checkbox',
                    }} disabled color="primary" />}
                    label="ALTER"
                    labelPlacement="top"
                />
            </Grid>
            <Grid item xs={3}>
                <FormControl className={classes.formControl}>
                    <InputLabel htmlFor="age-helper">Database Server</InputLabel>
                    <Select
                        value={values.age}
                        onChange={xhandleChange}
                        input={<Input name="age" id="age-helper" />}
                    >
                        <MenuItem value="">
                            <em>None</em>
                        </MenuItem>
                        <MenuItem value="db1.blah.com">db1.blah.com</MenuItem>
                        <MenuItem value="db2.blah.com">db2.blah.com</MenuItem>
                        <MenuItem value="db3.blah.com">db3.blah.com</MenuItem>
                    </Select>
                    <FormHelperText>Select the database for running the query</FormHelperText>
                </FormControl>
            </Grid>
            <Grid item xs={12}>
                <TextField
                    required
                    id="outlined-multiline-static"
                    label="SQL Query"
                    multiline
                    defaultValue={"-- SQL Query\n-- Paste it here\nSELECT * FROM XTABLE;"}
                    className={classes.textField}
                    margin="normal"
                    variant="outlined"
                    InputLabelProps={{
                        shrink: true,
                    }}
                    fullWidth
                />
            </Grid>
            <Grid item xs>
                <Button submit variant="contained" className={classes.button}>
                    PARSE
                </Button>
                <Button submit variant="contained" className={classes.button}>
                    SAVE
                </Button>
                <Button submit variant="contained" className={classes.button}>
                    APPROVE
                </Button>
                <Button submit variant="contained" className={classes.button}>
                    DISAPROVE
                </Button>
            </Grid>
        </form >
    );
}
