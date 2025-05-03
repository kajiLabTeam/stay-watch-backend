def time_to_minuts(time: str) -> int:
    """Convert time string to minutes"""
    h, m = map(int, time.split(":"))
    total = h * 60 + m
    return round(total / 30) * 30

def minuts_to_time(time: float) -> str:
    """Convert minutes to time string"""
    rounded = round(time / 30) * 30
    hours, minuts = divmod(int(rounded), 60)
    return f"{hours:02}:{minuts:02}"