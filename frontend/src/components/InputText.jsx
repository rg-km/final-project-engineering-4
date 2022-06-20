import { FormControl, FormLabel, Input } from '@chakra-ui/react';
import React from 'react';
import PropTypes from 'prop-types';

const InputText = ({value, name, label, handleChange, isRequired = false}) => {
    return <>
        <FormControl mb={[4]} isRequired={isRequired}>
            <FormLabel fontSize={["xs", "md"]}>
                {label}
            </FormLabel>
            <Input
                size={["md"]}
                name={name}
                value={value}
                onChange={handleChange}
                bg={'white'}
                isRequired={isRequired}
            />
        </FormControl>
    </>
}

InputText.propTypes = {
    value: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
    label: PropTypes.string,
    handleChange: PropTypes.func.isRequired,
    isRequired: PropTypes.bool
}

export default InputText;