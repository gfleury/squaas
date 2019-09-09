import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Stepper from '@material-ui/core/Stepper';
import Step from '@material-ui/core/Step';
import StepLabel from '@material-ui/core/StepLabel';
import Typography from '@material-ui/core/Typography';

function getSteps() {
    return ['Create query', 'Change query to ready', 'Approve/Disaprove query', 'Query approved', 'Query running', 'Done'];
}

export default class HorizontalLinearStepper extends React.Component {
    statusStep = {
        'new': 0,
        'pending': 1,
        'ready': 2,
        'approved': 3,
        'running': 4,
        'done': 5,
        'failed': 5,
    };


    constructor(props) {
        super(props);

        this.state = {
            status: props.status,
        }
        this.setState({ status: props.status });
        this.activeStep = 0;
        this.setActiveStep = 0;

    }

    steps = getSteps();

    classes = makeStyles(theme => ({
        button: {
            marginRight: theme.spacing(1),
        },
        instructions: {
            marginTop: theme.spacing(1),
            marginBottom: theme.spacing(1),
        },
    }));

    isStepOptional(step) {
        return step === 1;
    }

    render() {
        return (
            <Stepper activeStep={this.activeStep}>
                {this.steps.map((label, index) => {
                    const stepProps = {};
                    const labelProps = {};
                    if (this.isStepOptional(index)) {
                        labelProps.optional = <Typography variant="caption">Optional</Typography>;
                    }
                    if (index <= this.statusStep[this.state.status]) {
                        stepProps.completed = true;
                    }
                    console.log("Index: " + index + " Info: " + this.statusStep[this.state.status] + " " + this.state.status);
                    return (
                        <Step key={label} {...stepProps}>
                            <StepLabel {...labelProps}>{label}</StepLabel>
                        </Step>
                    );
                })}
            </Stepper>
        );
    }
}