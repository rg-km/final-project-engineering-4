import { Button, FormControl, FormLabel, Input, InputGroup, InputRightElement } from '@chakra-ui/react';
import React, { useState } from 'react';
import PropTypes from 'prop-types';

const InputPassword = ({ value, name, label, handleChange, isRequired }) => {
    const [show, setShow] = useState(false);

    return <>
        <FormControl mb={[4]} isRequired={isRequired}>
            <FormLabel fontSize={["xs", "md"]}>
                {label}
            </FormLabel>
            <InputGroup>
                <Input
                    name={name}
                    value={value}
                    type={show ? "text" : "password"}
                    onChange={handleChange}
                    required
                    bg={'white'}
                />
                <InputRightElement width="4.5rem">
                    <Button h="1.75rem" size="sm" onClick={() => setShow(!show)}>
                        {show ? "Hide" : "Show"}
                    </Button>
                </InputRightElement>
            </InputGroup>
        </FormControl>
    </>
}

InputPassword.propTypes = {
    value: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
    label: PropTypes.string.isRequired,
    handleChange: PropTypes.func.isRequired,
    isRequired: PropTypes.bool.isRequired
}

export default InputPassword;