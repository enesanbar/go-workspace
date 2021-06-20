import React from 'react';
import ReactDOM from 'react-dom';
import CommentDetail from './CommentDetail';
import ApprovalCard from "./ApprovalCard";
import faker from "faker";

class App extends React.Component {

    render() {
        return (
            <div className="ui container comments">
                <ApprovalCard>
                    <CommentDetail
                        name={faker.name.firstName()}
                        weekday={faker.date.weekday()}
                        comment={faker.lorem.text()}
                        avatar={faker.image.avatar()}
                    />
                </ApprovalCard>

                <CommentDetail
                    name={faker.name.firstName()}
                    weekday={faker.date.weekday()}
                    comment={faker.lorem.sentence()}
                    avatar={faker.image.avatar()}
                />

                <CommentDetail
                    name={faker.name.firstName()}
                    weekday={faker.date.weekday()}
                    comment={faker.lorem.sentence()}
                    avatar={faker.image.avatar()}
                />

                <ApprovalCard>
                    <CommentDetail
                        name={faker.name.firstName()}
                        weekday={faker.date.weekday()}
                        comment={faker.lorem.text()}
                        avatar={faker.image.avatar()}
                    />
                </ApprovalCard>
            </div>
        );
    }
}

ReactDOM.render(<App/>, document.querySelector('#root'));
