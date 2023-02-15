import {observer} from "mobx-react-lite";
import {DatePicker, Space, Input, Row, Col, Button} from 'antd';
import {useFormState, useRouteFormSubmit} from './hooks';
import styles from './RouteForm.module.css';
import dayjs from "dayjs";


export type RouteFormData = {
    from: string
    to: string
    date: string
}


const RouteForm = observer(function RouteForm() {
    const onSubmit = useRouteFormSubmit()
    const [form, setForm] = useFormState();

    const onFieldChange = (name: string, value: string) => setForm({...form, [name]: value});
    return (
        <Space>
            <Input.Group size='large'>
                <Space direction='horizontal'>
                    <Input value={form.from} placeholder='Откуда' size='large' onChange={evt => onFieldChange('from', evt.target.value)}/>
                    <div className={styles.wrapper}>
                        <span style={{whiteSpace: "nowrap"}}>&lt;-&gt;</span>
                    </div>
                    <Input value={form.to} placeholder='Куда' size='large' onChange={evt => onFieldChange('to', evt.target.value)}/>
                    <DatePicker value={dayjs(form.date)} placeholder='Дата' size='large' onChange={(dayjs, str) => onFieldChange('date', str)}/>
                    <Button size='middle' type='primary' onClick={() => onSubmit(form)}>Найти</Button>
                </Space>
            </Input.Group>
        </Space>
    );
});

export default RouteForm;