fn main() {
	let number_list = vec![34,24,23,12,64];
	let result = largest(&number_list);
	println!("result is {}", result);
	let char_list = vec!['y', 'm', 's', 'a'];
	let result = largest(&char_list);
	println!("result is {}", result);
}

fn largest<T: PartialOrd> (list: &[T]) -> &T {
	let mut max = &list[0];
    
    for i in list {
    	if i > max {
        	max = i;
        }
    }
    
    max
    
}
