import React, { ReactElement } from 'react';
import { Card, CardBody } from '@patternfly/react-core';

type BuildtimeMessageProps = {
    message: string;
};

function BuildtimeMessage({ message }: BuildtimeMessageProps): ReactElement {
    return (
        <Card isFlat className="pf-u-mb-md">
            <CardBody>{message}</CardBody>
        </Card>
    );
}

type BuildtimeMessagesProps = {
    violations?: {
        message: string;
    }[];
};

function BuildtimeMessages({ violations = [] }: BuildtimeMessagesProps): ReactElement {
    return (
        <>
            {violations.map(({ message }) => (
                <BuildtimeMessages key={message} message={message} />
            ))}
        </>
    );
}

export default BuildtimeMessages;
