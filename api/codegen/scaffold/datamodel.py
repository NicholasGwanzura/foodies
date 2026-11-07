import re,json,collections
def parse(path='docs/data-model.md'):
    L=open(path).read().split('\n'); T={}; order=[]; cur=None; sect=None; i=0
    while i<len(L):
        l=L[i]
        if l.startswith('### ') or l.startswith('## '): sect=l.lstrip('# ').strip()
        m=re.match(r'#### `(\w+)`',l); c=re.match(r'\*\*`(\w+)`\*\* \(child of `(\w+)`\)\.?\s*(.*)',l)
        m2=re.match(r'\*\*`(\w+)`\*\*',l)
        if m or c or m2:
            name=(m or c or m2).group(1)
            cur=dict(name=name,parent=c.group(2) if c else None,section=sect,cols=[],status=[],prefix=None,phase=None,restricted=False,desc='')
            T[name]=cur; order.append(name)
        elif cur is not None:
            mm=re.match(r'Phase (\d)\.(.*)',l)
            if mm:
                cur['phase']=int(mm.group(1))
                p=re.search(r'Reference prefix `(\w+)`',l)
                if p: cur['prefix']=p.group(1)
                cur['restricted']='Restricted' in l
            mm=re.match(r'\| `(\w+)` \| ([^|]+) \| ([^|]*) \| (.*) \|$',l)
            if mm: cur['cols'].append(dict(name=mm.group(1),type=mm.group(2).strip(),req=mm.group(3).strip()=='yes',notes=mm.group(4)))
            if l.startswith('Status values:'): cur['status']=re.findall(r'`(\w+)`',l)
        i+=1
    return T,order
if __name__=='__main__':
    T,o=parse(); print(len(T)); c=collections.Counter(x['type'] for t in T.values() for x in t['cols']); print(c)
    print([ (n,t['parent']) for n,t in T.items() if t['phase'] is None and not t['parent']][:20])
    print(collections.Counter(re.sub(r'`\w+`','`x`',x['notes'])[:40] for t in T.values() for x in t['cols']).most_common(25))
    fk=collections.Counter(m for t in T.values() for x in t['cols'] for m in re.findall(r'FK `(\w+)(?:\.\w+)?`',x['notes'])); print([k for k in fk if k not in T])
    print([ (t['name'],x['name'],x['type'],x['notes']) for t in T.values() for x in t['cols'] if x['type'] in('time','json[]','enum[]','text[]') or x['type'].endswith('[]')][:30])
