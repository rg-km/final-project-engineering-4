import { Box, Image, Text } from '@chakra-ui/react';
import React from 'react';
import { Link } from 'react-router-dom';
import PropTypes from 'prop-types'

const RoleItem = ({redirect, img, title, bg}) => {
    return <Box>
        <Link to={redirect}>
            <Box borderRadius={['lg']} boxShadow={['md']} p={[2, 4]} mb={[2]} bg={bg}>
                <Image src={img} alt={'gambar'} m={"auto"} />
            </Box>
            <Text textAlign={'center'} fontSize={['sm', 'md']}>{title}</Text>
        </Link>
    </Box>
}

RoleItem.propTypes = {
    redirect: PropTypes.string.isRequired,
    img: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    bg: PropTypes.string.isRequired
}

export default RoleItem;