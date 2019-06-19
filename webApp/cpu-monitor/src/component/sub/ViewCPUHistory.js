import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import RTChart from 'react-rt-chart';

class ViewCpu extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isApiFetched: false,
            data: {
                 date: new Date(),
                Usage: 0,
                Temp: 0
            }
        };
    }

    componentDidMount() {
        this.interval = setInterval(() => {fetch('http://localhost:82/api/v1/cpu/state/current ')
            .then(res => res.json())
            .then(json => {
                this.setState({
                    isApiFetched: true,
                    data: {date: new Date(),
                    Usage: json.Usage,
                    Temp: json.Temp
                    }
                }
                )
            });
            this.forceUpdate()}
            , 1000);
    }
    componentWillUnmount() {
        clearInterval(this.interval);
    }

    render() {

        var chart = {
            axis: {
                y: { min: 0, max: 100 }
            },
            point: {
                show: false
            }
    };
        return (

            <Container fluid={true}>
                <Row style={{ height: '60px' }}>
                    <Col lg={12}>
                        <p style={{ margin: '0px', fontSize: '30px', color: 'grey' }}>History</p>
                    </Col>
                </Row>
                <Row>
                    <Col lg={6}>
                        <div
                            style={{
                                width: "100%",
                                height: "300px"
                            }}
                        > 
                            <RTChart
                                chart={chart}
                                fields={['Usage','Temp']}
                                data={this.state.data} />
                        }
                            
                        </div>
                    </Col>
                </Row>
            </Container>
        );
    }
}

export default ViewCpu; // Don’t forget to use export default!