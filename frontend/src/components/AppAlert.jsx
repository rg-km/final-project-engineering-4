import { Alert, AlertIcon } from '@chakra-ui/react';
import React, { useCallback, useEffect } from 'react';
import PropTypes from 'prop-types';

const AppAlert = ({ alert, setAlert }) => {
    const changeAlert = useCallback(() => {
        setAlert({
            status: false,
            type: "warning",
            message: ""
        })
    }, [setAlert])

    useEffect(() => {
        if (alert.status) {
            setTimeout(() => {
                changeAlert()
            }, 3000);
        }
    }, [alert, changeAlert])

    return <>
        {
            alert.status && (
                <Alert status={alert.type} fontSize={['xs', 'md']}>
                    <AlertIcon />
                    {alert.message}
                </Alert>
            )
        }
    </>
}

AppAlert.propTypes = {
    alert: PropTypes.object.isRequired,
    setAlert: PropTypes.func.isRequired
}

export default AppAlert;