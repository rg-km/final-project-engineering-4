import { Link } from 'react-router-dom';

const { Box, Image } = require('@chakra-ui/react')

const AddClass = ({to, onClick}) => {
    return <Box
        display={'flex'}
        bg={'white'}
        p={{ "base": 2, "md": 4 }}
        as={Link}
        onClick={onClick}
        to={to}
        rounded={'md'}
        shadow={'sm'}
        border={'1px'}
        borderColor={'transparent'}
        transition={'all 200ms ease'}
        _hover={{ borderColor: 'primary.500' }}>
        <Image src={require('@images/icons/plus.png')} alt={'plus icon'} m={'auto'} w={'100%'} />
    </Box>
}

export default AddClass;