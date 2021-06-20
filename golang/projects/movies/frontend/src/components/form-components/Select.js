const Select = (props) => {
    return (
        <div className="mb-3">
            <label htmlFor={props.name} className="form-label">
                {" "}
                { props.title }{" "}
            </label>
            <select className="form-control" name={props.name} id={props.name} value={props.value}
                    onChange={props.handleChange}>
                <option value="" className="form-select">{props.placeholder}</option>
                {props.options.map((option) => {
                    return (
                        <option key={option.id} value={option.id} label={option.value} className="form-select">{option.value}</option>
                    )
                })}
            </select>
        </div>
    )
}

export default Select;
