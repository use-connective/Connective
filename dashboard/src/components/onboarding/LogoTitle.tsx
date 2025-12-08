interface LogoAndTitleProps {
    Logo: string;
    Width: number;
}

function LogoAndTitle(props: LogoAndTitleProps) {
    return (
        <>
            {/* Logo + Title */}
            <div className="flex items-center space-x-2">
                <img src={props.Logo} width={props.Width} alt="Logo" />
                <p className="text-md font-medium lg:text-2xl lg:font-medium text-gray-800">
                    Connective
                </p>
            </div>
        </>
    );
}

export default LogoAndTitle;
